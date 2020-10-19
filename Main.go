package main

import (
    "fmt"
    "log"
	"net/http"
	"time"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strconv"
)
type Person struct {
    Name string
	Email string
	RSVP string
	Meetings []int
}
type Meeting struct {
	Id int
	Title string
	Title1 int
	Title2 int
	Participants []Person
	Timestamp time.Time
}


func handler(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tej:tpa4401@first-bvv78.gcp.mongodb.net/Appointy?retryWrites=true&w=majority"))
if err != nil {
	log.Fatal(err)
}
ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
err = client.Connect(ctx)
if err != nil {
	log.Fatal(err)
}
defer client.Disconnect(ctx)
err = client.Ping(ctx,readpref.Primary())
if err != nil {
	log.Fatal(err)
}
collection := client.Database("Appointy").Collection("meetings")
collection1 := client.Database("Appointy").Collection("persons")
	switch r.Method {
		case "GET":
			id := r.FormValue("id")
			start := r.FormValue("start")
			end := r.FormValue("end")
			participant := r.FormValue("participant")
			if(id != ""){
				int1, err := strconv.ParseInt(id, 6, 12)
				ctx,_ = context.WithTimeout(context.Background(), 30*time.Second)
				cur, err := collection.Find(ctx,bson.M{"Id": int1} )
				if err != nil { log.Fatal(err) }
				defer cur.Close(ctx)
				for cur.Next(ctx) {
				var result Meeting
				err := cur.Decode(&result)
				if err != nil { log.Fatal(err) }
				fmt.Println("id is " + 	strconv.Itoa(result.Id))
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(result)
				}
				if err := cur.Err(); err != nil {
				log.Fatal(err)
				}
			}
			if(start != "" && end != ""){
				ctx,_ = context.WithTimeout(context.Background(), 30*time.Second)
				cur, err := collection.Find(ctx,bson.D{{}} )
				if err != nil { log.Fatal(err) }
				defer cur.Close(ctx)
				for cur.Next(ctx) {
				var result Meeting
				err := cur.Decode(&result)
				if err != nil { log.Fatal(err) }
				int1, err := strconv.Atoi(start)
				int2, err := strconv.Atoi(end)
				fmt.Println("id is " + 	strconv.Itoa(result.Title1) + " " + strconv.Itoa(result.Title2) + " " + strconv.Itoa(int(int1)) + " " + strconv.Itoa(int(int2)))
				if(result.Title1 >= int(int1) && result.Title2 <= int(int2)){
				fmt.Println("success")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(result)
				}
				}
				if err := cur.Err(); err != nil {
				log.Fatal(err)
				}
			}
			if(participant != ""){
				ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
				cur, err := collection1.Find(ctx, bson.M{"email":participant})
				if err != nil { log.Fatal(err) }
				defer cur.Close(ctx)
				for cur.Next(ctx) {
				var result Person
				err := cur.Decode(&result)
				if err != nil { log.Fatal(err) }
				for i := 0; i < len(result.Meetings); i++ {
					ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
				cur, err := collection.Find(ctx, bson.M{"Id":result.Meetings[i]})
				if err != nil { log.Fatal(err) }
				defer cur.Close(ctx)
				for cur.Next(ctx) {
				var res Meeting
				err := cur.Decode(&res)
				if err != nil { log.Fatal(err) }
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(res)
				}
			}
				if err := cur.Err(); err != nil {
				log.Fatal(err)
				}
		}
	}
		case "POST":
			var meeting Meeting
			err := json.NewDecoder(r.Body).Decode(&meeting)
			if err != nil { log.Fatal(err) }
			Id := meeting.Id
			Title := meeting.Title
			Participants := meeting.Participants
			Title1:= meeting.Title1
			Title2 := meeting.Title2
			timestamp := time.Now()
			meeting.Timestamp = time.Now()
			// q := r.FormValue("q")
			// fmt.Println(r.Body)
			fmt.Println(meeting)
			ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
			res, err := collection.InsertOne(ctx, bson.M{"Id":Id,
			"Title": Title,
			"Participants": Participants,
			"Title1":Title1,
			"Title2": Title2,
			"timestamp": timestamp})
			id := res.InsertedID
			fmt.Print(id)
			for i := 0; i < len(meeting.Participants); i++ {
			// 	ctx,_ = context.WithTimeout(context.Background(), 30*time.Second)
			// cur, err := collection.Find(ctx,bson.M{"Name": meeting.Participants[i].Name,"Email":meeting.Participants[i].Email,"RSVP":meeting.Participants[i].RSVP} )
			// if err != nil { log.Fatal(err) }
			// defer cur.Close(ctx)
			// for cur.Next(ctx) {
			//    var result bson.M
			//    err := cur.Decode(&result)
			//    if err != nil { log.Fatal(err) }
			//    do something with result....
			// }
			// if err := cur.Err(); err != nil {
			//   log.Fatal(err)
			// }
				person := Person{meeting.Participants[i].Name,meeting.Participants[i].Email,meeting.Participants[i].RSVP,[]int{meeting.Id}}
				result, _ := collection1.InsertOne(ctx, person)
				id1 := result.InsertedID
				fmt.Print(id1)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(meeting)
default:
	fmt.Println("Sorry, only GET and POST methods are supported.")
}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/meeting", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
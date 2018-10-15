package handler

import (
	"context"
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/Eric-GreenComb/recharge-linked/bean"
	"github.com/Eric-GreenComb/recharge-linked/config"
)

// Alter Alter
func Alter(c *gin.Context) {
	conn, err := grpc.Dial(config.Dgraph.Host, grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	op := &api.Operation{}
	op.Schema = `
		name: string @index(exact) .
		age: int .
		married: bool .
		loc: geo .
		dob: datetime .
	`

	ctx := context.Background()
	err = dg.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "OK"})
}

// Mutate Mutate
func Mutate(c *gin.Context) {

	conn, err := grpc.Dial(config.Dgraph.Host, grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	dob := time.Date(1980, 01, 01, 23, 0, 0, 0, time.UTC)
	// While setting an object if a struct has a Uid then its properties in the graph are updated
	// else a new node is created.
	// In the example below new nodes for Alice, Bob and Charlie and school are created (since they
	// dont have a Uid).
	p := bean.Person{
		Name:    "Tita",
		Age:     26,
		Married: true,
		Location: bean.Loc{
			Type:   "Point",
			Coords: []float64{1.1, 2},
		},
		Dob: &dob,
		Raw: []byte("raw_bytes"),
		Friends: []bean.Person{{
			Name: "Bob",
			Age:  24,
		}, {
			Name: "Eric",
			Age:  46,
		}, {
			Name: "Charlie",
			Age:  29,
		}},
		School: []bean.School{{
			Name: "Crown Public School",
		}},
	}

	ctx := context.Background()

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	_, err = dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "OK"})
}

// Query Query
func Query(c *gin.Context) {

	conn, err := grpc.Dial(config.Dgraph.Host, grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	defer conn.Close()

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	ctx := context.Background()

	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	// variables := map[string]string{"$id": assigned.Uids["blank-0"]}
	const q = `{
		person(func: has(name)) {
		  uid
		  name
		  dob
		  age
		  loc
		  raw_bytes
		  married
		  friend{
			  name
			  age
		  }
		  school {
			  name
		  }
		}
	  }`

	// resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	resp, err := dg.NewTxn().Query(ctx, q)
	if err != nil {
		log.Fatal(err)
	}

	// type Root struct {
	// 	Me []bean.Person `json:"me"`
	// }

	// var r Root
	// err = json.Unmarshal(resp.Json, &r)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Me: %+v\n", r.Me)
	// R.Me would be same as the person that we set above.

	// fmt.Println(string(resp.Json))
	// Output: {"me":[{"name":"Alice","dob":"1980-01-01T23:00:00Z","age":26,"loc":{"type":"Point","coordinates":[1.1,2]},"raw_bytes":"cmF3X2J5dGVz","married":true,"friend":[{"name":"Bob","age":24}],"school":[{"name":"Crown Public School"}]}]}

	c.String(http.StatusOK, string(resp.Json))
}

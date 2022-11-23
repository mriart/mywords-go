package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
func mix(in []string) []string {
	out := make([]string, 0)
	for i := 0; i < 3; i++ {
		out = append(out, in[i])
	}
	copy(out, in)
	out = append(out, "4")
	fmt.Println(out)
	return out
}
*/

func shuffle(in []string) {
	for i := range in {
		rand.Seed(time.Now().UnixMicro())
		j := rand.Intn(i + 1)
		in[i], in[j] = in[j], in[i]
	}
}

func pre3Append(pre []string, post []string) []string {
	post = append(post, "")
	post = append(post, "")
	post = append(post, "")
	copy(post[3:], post)
	for i := 0; i < 3; i++ {
		post[i] = pre[i]
	}
	return post
}

func main() {
	collection := []string{"1", "2", "3", "11", "22", "33"}
	pre3coll := []string{collection[0], collection[1], collection[2]}
	fmt.Println(collection)

	shuffle(collection)
	fmt.Println(collection)

	shuffle(collection)
	fmt.Println(collection)

	shuffle(collection)
	fmt.Println(collection)

	//collection = append(pre3coll, collection...)
	collection = pre3Append(pre3coll, collection)
	fmt.Println(collection)
}

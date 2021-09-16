package main

import (
	"log"
)

type middleware interface {
	execute(*req)
}

type mid struct {
	title string
	next  middleware
}

type req struct {
	header string
}

func (m *mid) execute(r *req) {
	log.Print(m.title, r.header)
	if m.next == nil {
		log.Print("middleware ended")
		return
	}
	m.next.execute(r)
}

func main() {
	r := req{}
	mid3 := mid{
		title: "first middleware",
		next:  nil,
	}
	mid2 := mid{
		title: "second middleware",
		next:  &mid3,
	}

	mid1 := mid{
		title: "third middleware",
		next:  &mid2,
	}
	mid1.execute(&r)
}

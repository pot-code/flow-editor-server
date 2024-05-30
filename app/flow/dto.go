package flow

import "time"

type FlowListItem struct {
	Id        int
	Title     string
	CreatedAt time.Time
}

type FlowDetail struct {
	Id        int
	Title     string
	Nodes     string
	Edges     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

package model

type Group struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	MemberIDs []int  `json:"member_ids"`
}

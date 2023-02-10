package models

type Domain struct {
	ID     int    `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Status int    `db:"status" json:"status"`
}

func NewDomain() *Domain {
	return &Domain{}
}

func NewDomains() *[]Domain {
	return &[]Domain{}
}

func (d *Domain) SetID(id int) {
	d.ID = id
}

func (d *Domain) SetName(domainName string) {
	d.Name = domainName
}

func (d *Domain) SetStatus(status int) {
	d.Status = status
}

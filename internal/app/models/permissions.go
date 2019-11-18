package models

// Permissions contains the moderator's permissions
// of a moderator user in the control panel
type Permissions struct {
	Users    bool `bson:"users" json:"users"`
	Events   bool `bson:"events" json:"events"`
	Tickets  bool `bson:"tickets" json:"tickets"`
	Sponsors bool `bson:"sponsors" json:"sponsors"`
}

// NewPermissions is a permission factory
func NewPermissions(users bool, events bool, tickets bool, sponsors bool) *Permissions {
	return &Permissions{
		Users:    users,
		Tickets:  tickets,
		Events:   events,
		Sponsors: sponsors,
	}
}

// GetUsers returns moderator's permissions for users
func (p *Permissions) GetUsers() bool {
	return p.Users
}

// SetUsers changes moderator's permissions for users
func (p *Permissions) SetUsers(users bool) {
	p.Users = users
}

// GetEvents returns moderator's permissions for events
func (p *Permissions) GetEvents() bool {
	return p.Events
}

// SetEvents changes moderator's permissions for events
func (p *Permissions) SetEvents(events bool) {
	p.Events = events
}

// GetTickets returns moderator's permission for tickets
func (p *Permissions) GetTickets() bool {
	return p.Tickets
}

// SetTickets changes moderator's permissions for tickets
func (p *Permissions) SetTickets(tickets bool) {
	p.Tickets = tickets
}

// GetSponsors returns moderator's permissions for sponsors
func (p *Permissions) GetSponsors() bool {
	return p.Sponsors
}

// SetSponsors changes moderator's permissions for sponsors
func (p *Permissions) SetSponsors(sponsors bool) {
	p.Sponsors = sponsors
}

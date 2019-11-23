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

// hasUserPermissions returns moderator's permissions for users
func (p *Permissions) hasUserPermissions() bool {
	return p.Users
}

// SetUserPermissions changes moderator's permissions for users
func (p *Permissions) SetUserPermissions(users bool) {
	p.Users = users
}

// hasEventPermissions returns moderator's permissions for events
func (p *Permissions) hasEventPermissions() bool {
	return p.Events
}

// SetEventPermissions changes moderator's permissions for events
func (p *Permissions) SetEventPermissions(events bool) {
	p.Events = events
}

// hasTicketPermissions returns moderator's permission for tickets
func (p *Permissions) hasTicketPermissions() bool {
	return p.Tickets
}

// SetTicketPermisiions changes moderator's permissions for tickets
func (p *Permissions) SetTicketPermisiions(tickets bool) {
	p.Tickets = tickets
}

// hasSponsorPermissions returns moderator's permissions for sponsors
func (p *Permissions) hasSponsorPermissions() bool {
	return p.Sponsors
}

// SetSponsorPermissions changes moderator's permissions for sponsors
func (p *Permissions) SetSponsorPermissions(sponsors bool) {
	p.Sponsors = sponsors
}

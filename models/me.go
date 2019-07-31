package models

func (me *Me) GetId() string        { return me.Id }
func (me *Me) GetName() string      { return me.Name }
func (me *Me) GetAuthor() string    { return me.Name }
func (me *Me) GetParentId() string  { return me.Id }
func (me *Me) GetTitle() string     { return me.Name }
func (me *Me) GetBody() string      { return me.Name }
func (me *Me) GetKarma() float64    { return me.CommentKarma + me.LinkKarma }
func (me *Me) GetUps() float64      { return 0 }
func (me *Me) GetDowns() float64    { return 0 }
func (me *Me) GetSubreddit() string { return me.Name }
func (me *Me) GetCreated() float64  { return me.Created }
func (me *Me) GetFlair() string     { return "" }
func (me *Me) GetUrl() string       { return me.IconImg }
func (me *Me) IsRoot() bool         { return true }

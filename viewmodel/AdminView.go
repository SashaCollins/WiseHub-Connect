package viewmodel

import (
	"github/SashaCollins/Wisehub-Connect/model/data"
	"github/SashaCollins/Wisehub-Connect/model/listener"
)

type AdminView struct {
	ds data.Datastore
	l listener.Listener
}
func (av *AdminView) Show() {

}
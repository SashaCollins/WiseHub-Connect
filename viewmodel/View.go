/*
Interface for Views.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
*/
package viewmodel

import (
	_ "github/SashaCollins/Wisehub-Connect/model/data"
	_ "github/SashaCollins/Wisehub-Connect/model/listener"
)

type View interface {
	Show()
}
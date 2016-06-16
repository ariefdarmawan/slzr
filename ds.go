package slzr

import (
    "github.com/eaciit/dbox"
)

type DataSource struct{
    ID string
    Connection *dbox.ConnectionInfo
}
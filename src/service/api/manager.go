package api

//go:generate mockery -name Manager -outpkg apimocks -output ./apimocks -dir .
type Manager interface {
}


// impl

type manager struct {
	stg   string
}
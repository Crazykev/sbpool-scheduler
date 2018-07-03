package types

type Machine struct {
	// ID is id of machine, parsed from machine_xxx
	ID int
	// Disk is disk type resource
	Disk GenericResource
	// P is P type resource
	P GenericResource
	// M is M type resource
	M GenericResource
	// PM is PM type resource
	PM GenericResource
	// Instances are all instances running on current Machine
	Instances []*Instance
}

type App struct {
	// ID is id of app, parsed from app_xxx
	ID int
}

type SequenceResource struct {
	GenericResource
	UtilizationSequence []int
	MaxUtilization      int
	MinUtilization      int
}

func (sr *SequenceResource) Allocate(requestSeq []int) error {
	if len(requestSeq) != len(sr.UtilizationSequence) {
		return error.Error("length of request sequence is not equal to resource utilization sequence")
	}
	newUtil := []int{}
	for index := 1; index < len(requestSeq); index++ {
		cur := requestSeq[index] + sr.UtilizationSequence
		if cur > sr.Allocatable {
			return error.Error("request sequence exceed allocatable in timestamp %d", index)
		}
	}
	return nil
}

// GenericResourceType is a enum type of normal integer type resource
type GenericResourceType int

const (
	Disk GenericResourceType = iota
	P
	M
	PM
)

// GenericResourceType is normal integer type resource
type GenericResource struct {
	Capacity    int
	Allocatable int
	Requested   int
}

func NewGeneraicResource(capacity int) GenericResource {
	return GenericResource{
		Capacity:    capacity,
		Allocatable: capacity,
	}
}

func (g *GenericResource) Allocate(request int) error {
	if request < 0 {
		return error.Error("resource request should be positive")
	}
	if request+g.Requested > g.Allocatable {
		return error.Error("resource request exceed allocatable resource")
	}
	g.Requested += request
	return nil
}

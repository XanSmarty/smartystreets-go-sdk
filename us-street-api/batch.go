package us_street

// Batch stores input records and settings related to a group of addresses to be verified in a batch.
type Batch struct {
	lookups []*Lookup

	standardizeOnly bool
	includeInvalid  bool
}

// NewBatch creates a new, empty batch.
func NewBatch() *Batch {
	return &Batch{}
}

// StandardizeOnly sets the X-Standardize-Only header value.
func (b *Batch) StandardizeOnly(on bool) {
	b.standardizeOnly = on
}

// IncludeInvalid sets the X-Include-Invalid header value.
func (b *Batch) IncludeInvalid(on bool) {
	b.includeInvalid = on
}

// Append includes the record in the collection to be sent if there is still room (max: 100).
func (b *Batch) Append(record *Lookup) bool {
	hasSpace := len(b.lookups) < MaxBatchSize
	if hasSpace {
		b.lookups = append(b.lookups, record)
	}
	return hasSpace
}

func (b *Batch) attach(candidate *Candidate) {
	i := candidate.InputIndex
	b.lookups[i].Results = append(b.lookups[i].Results, candidate)
}

// IsFull returns true when the batch has 100 lookups, false in every other case.
func (b *Batch) IsFull() bool {
	return b.Length() == MaxBatchSize
}

func (b *Batch) isEmpty() bool {
	return b.Length() == 0
}

// Length returns
func (b *Batch) Length() int {
	return len(b.lookups)
}

// Records returns the internal records collection.
func (b *Batch) Records() []*Lookup {
	return b.lookups
}

// Clear clears the internal collection.
func (b *Batch) Clear() {
	b.lookups = nil
}

// Reset clears the internal collection and resets settings.
func (b *Batch) Reset() {
	b.Clear()
	b.standardizeOnly = false
	b.includeInvalid = false
}

const MaxBatchSize = 100
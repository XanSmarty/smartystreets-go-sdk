package us_street

import (
	"encoding/json"
	"strconv"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type BatchFixture struct {
	*gunit.Fixture
}

func (f *BatchFixture) TestCapacityIsLimitedAt100Inputs() {
	batch := NewBatch()

	f.So(batch.Length(), should.Equal, 0)
	f.So(batch.Records(), should.HaveLength, 0)

	for x := 0; x < 100; x++ {
		f.So(batch.Append(&Input{InputID: strconv.Itoa(x)}), should.BeTrue)
	}
	f.So(batch.Length(), should.Equal, 100)
	f.So(batch.Records(), should.HaveLength, 100)

	for x := 100; x < 200; x++ {
		f.So(batch.Append(&Input{InputID: strconv.Itoa(x)}), should.BeFalse)
	}

	f.So(batch.Length(), should.Equal, 100)
	f.So(batch.Records(), should.HaveLength, 100)
}

func (f *BatchFixture) TestJSONSerializationShouldNeverFail() {
	batch := NewBatch()
	batch.Append(&Input{
		Street:        "This",
		Street2:       "test",
		Secondary:     "exists",
		City:          "to",
		State:         "ensure",
		ZIPCode:       "the",
		LastLine:      "input",
		Addressee:     "always",
		Urbanization:  "serializes",
		InputID:       "successfully",
		MaxCandidates: 7,
	})
	serialized, err := json.Marshal(batch.records)
	f.So(err, should.BeNil)
	f.So(serialized, should.NotBeEmpty)
}

func (f *BatchFixture) TestClearRemovesAllRecords() {
	batch := NewBatch()
	batch.StandardizeOnly(true)
	batch.IncludeInvalid(true)

	for x := 0; x < 100; x++ {
		f.So(batch.Append(&Input{InputID: strconv.Itoa(x)}), should.BeTrue)
	}

	batch.Clear()

	f.So(batch.Length(), should.Equal, 0)
	f.So(batch.standardizeOnly, should.BeTrue)
	f.So(batch.includeInvalid, should.BeTrue)
}

func (f *BatchFixture) TestResetRemovesAllRecordsAndResetsSettings() {
	batch := NewBatch()
	batch.StandardizeOnly(true)
	batch.IncludeInvalid(true)

	for x := 0; x < 100; x++ {
		f.So(batch.Append(&Input{InputID: strconv.Itoa(x)}), should.BeTrue)
	}

	batch.Reset()

	f.So(batch.Length(), should.Equal, 0)
	f.So(batch.standardizeOnly, should.BeFalse)
	f.So(batch.includeInvalid, should.BeFalse)
}

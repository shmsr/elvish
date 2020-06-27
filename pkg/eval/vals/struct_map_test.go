package vals

import (
	"testing"

	"github.com/xiaq/persistent/hash"
)

type testStructMap struct {
	Name        string
	ScoreNumber float64
}

func (testStructMap) IsStructMap() {}

// Structurally identical to testStructMap.
type testStructMap2 struct {
	Name        string
	ScoreNumber float64
}

func (testStructMap2) IsStructMap() {}

type testStructMap3 struct {
	Name  string
	score float64
}

func (testStructMap3) IsStructMap() {}

func (m testStructMap3) Score() float64 {
	return m.score + 10
}

func TestStructMap(t *testing.T) {
	TestValue(t, testStructMap{}).
		Kind("structmap").
		Bool(true).
		Hash(hash.DJB(Hash(""), Hash(0.0))).
		Repr(`[&name='' &score-number=(float64 0)]`).
		Len(2).
		Equal(testStructMap{}).
		NotEqual("a", MakeMap(), testStructMap{"a", 1.0}).
		// StructMap's are nominally typed. This may change in future.
		NotEqual(testStructMap2{}).
		HasKey("name", "score-number").
		HasNoKey("bad", 1.0).
		IndexError("bad", NoSuchKey("bad")).
		IndexError(1.0, NoSuchKey(1.0)).
		AllKeys("name", "score-number").
		Index("name", "").
		Index("score-number", 0.0)

	TestValue(t, testStructMap{"a", 1.0}).
		Kind("structmap").
		Bool(true).
		Hash(hash.DJB(Hash("a"), Hash(1.0))).
		Repr(`[&name=a &score-number=(float64 1)]`).
		Len(2).
		Equal(testStructMap{"a", 1.0}).
		NotEqual(
			"a", MakeMap("name", "", "score-number", 1.0),
			testStructMap{}, testStructMap{"a", 2.0}, testStructMap{"b", 1.0}).
		// Keys are tested above, thus omitted here.
		Index("name", "a").
		Index("score-number", 1.0)

	TestValue(t, testStructMap3{"a", 1.0}).
		Kind("structmap").
		Bool(true).
		Hash(hash.DJB(Hash("a"), Hash(11.0))).
		Repr(`[&name=a &score=(float64 11)]`).
		Len(2).
		Equal(testStructMap3{"a", 1.0}).
		NotEqual(
			"a", MakeMap("name", "", "score-number", 1.0),
			testStructMap{}, testStructMap{"a", 11.0}).
		// Keys are tested above, thus omitted here.
		Index("name", "a").
		Index("score", 11.0)
}

package gopinyin_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/caiguanhao/gopinyin"
)

func get1st(first interface{}, args ...interface{}) interface{} {
	return first
}

func assert(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%s should equal to %s", actual, expected)
	}
}

func TestSplit(t *testing.T) {
	assert(t, gopinyin.Split("cgh"), gopinyin.Pinyins{"c", "g", "h"})
	assert(t, gopinyin.Split("CGH"), gopinyin.Pinyins{"c", "g", "h"})
	assert(t, gopinyin.Split("caigh"), gopinyin.Pinyins{"cai", "g", "h"})
	assert(t, gopinyin.Split("caiguh"), gopinyin.Pinyins{"cai", "gu", "h"})
	assert(t, gopinyin.Split("zhongguo"), gopinyin.Pinyins{"zhong", "guo"})
}

func TestAbbr(t *testing.T) {
	assert(t, gopinyin.Split("zhoguo").Abbreviate(), gopinyin.Pinyins{"z", "g"})
}

func TestExpand(t *testing.T) {
	assert(t, gopinyin.Split("zhoguo").Expand(), gopinyin.Pinyins{"zhong,zhou", "guo"})
}

func TestSQL(t *testing.T) {
	assert(t, gopinyin.Split("caiguanhao").Expand().SQL("pinyin"), "SEQUENCED_ARRAY_CONTAINS(pinyin, '{cai}', '{guan,guang}', '{hao}')")
	assert(t, gopinyin.Split("").Expand().SQL("pinyin"), "")
	assert(t, gopinyin.Split("123").Expand().SQL("pinyin"), "")
}

func TestValue(t *testing.T) {
	assert(t, get1st(gopinyin.Pinyins{}.Value()), "{}")
	assert(t, get1st(gopinyin.Pinyins{"0", "z", "1", "g", "2"}.Value()), "{z,g}")
	assert(t, get1st(gopinyin.Pinyins{"zhong", "guo"}.Value()), "{zhong,guo}")
	assert(t, get1st(gopinyin.Split("zg").Value()), "{z,g}")
}

func Example_postgreSQL() {
	fmt.Printf("SELECT * FROM data WHERE %s\n", gopinyin.Split("shouj").Expand().SQL("pinyin"))
	// Output:
	// SELECT * FROM data WHERE SEQUENCED_ARRAY_CONTAINS(pinyin, '{shou}', '{ji,jia,jian,jiang,jiao,jie,jin,jing,jiong,jiu,ju,juan,jue,jun}')
}

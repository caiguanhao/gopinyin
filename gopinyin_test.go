package gopinyin_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/caiguanhao/gopinyin"
)

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

func TestExpand(t *testing.T) {
	assert(t, gopinyin.Split("zhoguo").Expand(), gopinyin.Pinyins{"zhong,zhou", "guo"})
}

func TestSQL(t *testing.T) {
	assert(t, gopinyin.Split("caiguanhao").Expand().SQL("pinyin"), "pinyin && '{cai}' AND pinyin && '{guan,guang}' AND pinyin && '{hao}'")
	assert(t, gopinyin.Split("").Expand().SQL("pinyin"), "")
	assert(t, gopinyin.Split("123").Expand().SQL("pinyin"), "")
}

func Example_postgreSQL() {
	fmt.Printf("SELECT * FROM data WHERE %s\n", gopinyin.Split("shouj").Expand().SQL("pinyin"))
	// Output:
	// SELECT * FROM data WHERE pinyin && '{shou}' AND pinyin && '{ji,jia,jian,jiang,jiao,jie,jin,jing,jiong,jiu,ju,juan,jue,jun}'
}

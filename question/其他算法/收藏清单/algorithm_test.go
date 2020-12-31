package 收藏清单

import (
	"fmt"
	"testing"
)

/*
{{"leetcode","google","facebook"},{"google","microsoft"},{"google","facebook"},{"google"},{"amazon"}}

*/

/*
{{"arrtztkotazhufrsfczr","knzgidixqgtnahamebxf","zibvccaoayyihidztflj"},{"cffiqfviuwjowkppdajm","owqvnrhuzwqohquamvsz"},{"knzgidixqgtnahamebxf","owqvnrhuzwqohquamvsz","qzeqyrgnbplsrgqnplnl"},{"arrtztkotazhufrsfczr","cffiqfviuwjowkppdajm"},{"arrtztkotazhufrsfczr","knzgidixqgtnahamebxf","owqvnrhuzwqohquamvsz","qzeqyrgnbplsrgqnplnl","zibvccaoayyihidztflj"}}
*/
/*

	{"leetcode", "google", "facebook"}, {"google", "microsoft"}, {"google", "facebook"}, {"google"}, {"amazon"},
*/
var favoriteCompanies = [][]string{{"arrtztkotazhufrsfczr", "knzgidixqgtnahamebxf", "zibvccaoayyihidztflj"}, {"cffiqfviuwjowkppdajm", "owqvnrhuzwqohquamvsz"}, {"knzgidixqgtnahamebxf", "owqvnrhuzwqohquamvsz", "qzeqyrgnbplsrgqnplnl"}, {"arrtztkotazhufrsfczr", "cffiqfviuwjowkppdajm"}, {"arrtztkotazhufrsfczr", "knzgidixqgtnahamebxf", "owqvnrhuzwqohquamvsz", "qzeqyrgnbplsrgqnplnl", "zibvccaoayyihidztflj"}}

func TestSortArr(t *testing.T) {
	fmt.Println(peopleIndexes(favoriteCompanies))
}

func TestIsSubSet(t *testing.T) {
	fmt.Println(isSubSet(favoriteCompanies[2], favoriteCompanies[3]))
}

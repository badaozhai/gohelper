package gohelper

import (
	"testing"
	"fmt"
)

func TestGetValueByRegexp(t *testing.T) {
	str := `</html>
	<script xml:space="preserve">
	/*<![CDATA[*/
	 var hello=[{'231011','mobile_no':''}];
	 var world = 10;
	 var xx = 1;
	 var errorMsg = null;
	 /*]]>*/
	</script>
	<script id="passengerTemplate"`
	fmt.Println(GetValueByRegexp(str,`var hello=`,";"))
	fmt.Println(GetValueByRegexp(str,`var world = `,";"))
	fmt.Println(GetValueByRegexp(str,`var xx = `,";"))

	fmt.Println(GetValueBySplit(str,`var hello=`,";"))
	fmt.Println(GetValueBySplit(str,`var world = `,";"))
	fmt.Println(GetValueBySplit(str,`var xx = `,";"))
}
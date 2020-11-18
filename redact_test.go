package redact

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestSanitizeFixture(t *testing.T) {
	gunit.Run(new(SanitizeFixture), t)
}

type SanitizeFixture struct {
	*gunit.Fixture
}

//func (this *SanitizeFixture) TestRedactDOB() {
//	input := "Hello my name is John, my date of birth is 11/1/2000 and my employee's date of birth is 01-01-2001, oh also November 1, 2000, May 23, 2019, 23 June 1989, Sept 4, 2010."
//	expectedOutput := "Hello my name is John, my date of birth is [DOB REDACTED] and my employee's date of birth is [DOB REDACTED], oh also [DOB REDACTED], [DOB REDACTED], [DOB REDACTED], [DOB REDACTED]."
//
//	output := DateOfBirth(input) // TODO: change to ALL
//
//	this.So(output, should.Resemble, expectedOutput)
//}
//
//func (this *SanitizeFixture) TestRedactEmail() {
//	input := "Hello my name is John, my email address is john@test.com and my employee's email is jake@test.com and Jake Smith <jake@smith.com>."
//	expectedOutput := "Hello my name is John, my email address is [EMAIL REDACTED] and my employee's email is [EMAIL REDACTED] and Jake Smith <[EMAIL REDACTED]>."
//
//	output := Email(input)
//
//	this.So(output, should.Resemble, expectedOutput)
//}
//
//func (this *SanitizeFixture) SkipTestRedactCreditCard() {
//	input := "Hello my name is John, my Credit card number is: 1111-1111-1111-1111. My employees CC number is 1111111111111111 and 1111 1111 1111 1111 plus 1111111111111."
//	expectedOutput := "Hello my name is John, my Credit card number is: [CARD 1111****1111]. My employees CC number is [CARD 1111****1111] and [CARD 1111****1111] plus [CARD 1111****1111]."
//
//	output := CreditCard(input)
//
//	this.So(output, should.Resemble, expectedOutput)
//}

func (this *SanitizeFixture) TestMatchCreditCard() {
	input := "Blah 1234-5678-9012-3450. CC number is 1111111111118 and 1234 5678 9012 3450 1234-5678-9012-3455-230"
	this.So(matchCreditCard(input),should.Resemble, []match{{
		InputIndex: 5,
		Length:     19,
	}, {
		InputIndex: 39,
		Length:     13,
	}, {
		InputIndex: 57,
		Length:     19,
	}, {
		InputIndex: 77,
		Length:     23,
	}})
}

func(this *SanitizeFixture) TestRedactCreditCard(){
	input := "Blah 1234-5678-9012-3450. CC number is 1111111111118 and 1234 5678 9012 3450 1234-5678-9012-3450"
	expected := "Blah *******************. CC number is ************* and ******************* *******************"

	actual := All(input)

	this.So(actual, should.Equal,expected)
}

func (this *SanitizeFixture) TestMatchEmail(){
	input := "Blah test@gmail.com, our employee's email is test@gmail. and we have one more which may or not be an email " +
		"test@test."
	this.So(matchEmail(input), should.Resemble, []match{{
		InputIndex: 5,
		Length:     10,
	}, {
		InputIndex: 45,
		Length:     10,
	}, {
		InputIndex: 107,
		Length: 9,
	}})
}

func(this *SanitizeFixture) TestRedactEmail() {
	input := "Blah test@gmail.com, our employee's email is test@gmail. and we have one more which may or not be an email " +
		"test@test."
	expected := "Blah **********.com, our employee's email is **********. and we have one more which may or not be an email " +
		"*********."

	actual := All(input)

	this.So(actual, should.Equal,expected)
}

func(this *SanitizeFixture) TestMatchPhoneNum(){
	input := "Blah 801-111-1111"
	this.So(matchPhoneNum(input), should.Resemble, []match{{
		InputIndex: 5,
		Length: 12,
	}})
}
package main

// import (
// 	"github.com/dave/jennifer/jen"
// 	"github.com/montanaflynn/stats"
// )

// // structure for passing the testing data, expecteed results, and whether we expect a pass or fail
// type TestCase struct {
// 	name              string             // name of test case
// 	data              []stats.Coordinate //stats data from montanaflynn
// 	expectedSlope     float64            //m in y = mx + b
// 	expectedIntercept float64            //b in y = mx + b
// 	expectError       bool               //true if we expect an error
// }

// // creates unit test case
// // uses Jennifer library, so we can write multiple unit tests
// func generateTestCases(testCases []TestCase) {
// 	f := jen.NewFile("main") //package miain

// 	// setup the regression test
// 	f.Func().Id("TestRegress").Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(

// 		//declares the structure for the input data
// 		jen.Id("tests").Op(":=").Index().Struct(
// 			jen.Id("name").String(),
// 			jen.Id("data").Index().Qual("github.com/montanaflynn/stats", "Coordinate"),
// 			jen.Id("expectedSlope").Float64(),
// 			jen.Id("expectedIntercept").Float64(),
// 			jen.Id("expectError").Bool(),
// 		).ValuesFunc(func(g *jen.Group) {
// 			// pass throubgh the test cases
// 			for _, tc := range testCases {
// 				g.Values(
// 					jen.Lit(tc.name),
// 					jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").ValuesFunc(func(g *jen.Group) {
// 						for _, coord := range tc.data {
// 							g.Values(jen.Lit(coord.X), jen.Lit(coord.Y))
// 						}
// 					}),
// 					jen.Lit(tc.expectedSlope),
// 					jen.Lit(tc.expectedIntercept),
// 					jen.Lit(tc.expectError),
// 				)
// 			}
// 		}),
// 		// run the regression and pass the results
// 		jen.For(jen.List(jen.Id("_"), jen.Id("tc")).Op(":=").Range().Id("tests")).Block(
// 			jen.Id("t").Dot("Run").Call(jen.Id("tc").Dot("name"), jen.Func().Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
// 				jen.List(jen.Id("slope"), jen.Id("intercept"), jen.Id("err")).Op(":=").Id("regress").Call(jen.Id("tc").Dot("data")),
// 				jen.If(jen.Id("err").Op("!=").Nil().Op("&&").Id("tc").Dot("expectError").Op("==").False()).Block(
// 					jen.Id("t").Dot("Errorf").Call(jen.Lit("regress() returned an unexpected error: %v"), jen.Id("err")),
// 				),
// 				jen.If(jen.Id("err").Op("==").Nil().Op("&&").Id("tc").Dot("expectError")).Block(
// 					jen.Id("t").Dot("Errorf").Call(jen.Lit("regress() expected an error, got none")),
// 				),
// 				jen.If(jen.Id("slope").Op("!=").Id("tc").Dot("expectedSlope")).Block(
// 					jen.Id("t").Dot("Errorf").Call(jen.Lit("%s: Expected slope %v, got %v"), jen.Id("tc").Dot("name"), jen.Id("tc").Dot("expectedSlope"), jen.Id("slope")),
// 				),
// 				jen.If(jen.Id("intercept").Op("!=").Id("tc").Dot("expectedIntercept")).Block(
// 					jen.Id("t").Dot("Errorf").Call(jen.Lit("%s: Expected intercept %v, got %v"), jen.Id("tc").Dot("name"), jen.Id("tc").Dot("expectedIntercept"), jen.Id("intercept")),
// 				),
// 			)),
// 		),
// 	)

// 	// generate the test file
// 	err := f.Save("regression_test.go")
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func main() {

// 	// x. y data for three test cases
// 	data1 := []stats.Coordinate{
// 		{X: 1, Y: 3},
// 		{X: 2, Y: 5},
// 		{X: 3, Y: 7},
// 	}

// 	data2 := []stats.Coordinate{
// 		{X: 1, Y: 2},
// 		{X: 2, Y: 2},
// 		{X: 3, Y: 2},
// 	}

// 	data3 := []stats.Coordinate{
// 		{X: 1, Y: 6},
// 		{X: 2, Y: 8},
// 		{X: 3, Y: 10},
// 	}

// 	// populating the test case structure
// 	testCases := []TestCase{
// 		{
// 			name:              "Test Case 1: y = 2x + 1",
// 			data:              data1,
// 			expectedSlope:     2,
// 			expectedIntercept: 1,
// 			expectError:       false,
// 		},
// 		{
// 			name:              "Test Case 2: y = 0x + 2",
// 			data:              data2,
// 			expectedSlope:     0,
// 			expectedIntercept: 2,
// 			expectError:       false,
// 		},
// 		{
// 			name:              "Test Case 3: y = 2x + 4",
// 			data:              data3,
// 			expectedSlope:     2,
// 			expectedIntercept: 4,
// 			expectError:       false,
// 		},
// 	}

// 	generateTestCases(testCases)
// }

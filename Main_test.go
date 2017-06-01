package main_test

import (
        "testing"
        "github.com/stretchr/testify/assert"
        "github.com/go-test/deep"
        )

        type City struct{
          path1 string
          exit1 string
          cityName string
          path2 string
          exit2 string
        }


                func TestSomething(t *testing.T) {
                  assert := assert.New(t)

                  var a string = "Hello"
                  var b string = "Hello"

                  assert.Equal(a, b, "The two words should be the same.")
                }

        func TestCityType(t *testing.T) {

            City1 := City{
                path1 :     "A City Road",
                exit1 :     "A Exit Road",
                cityName :  "Napier",
                path2 :     "Another City Road",
                exit2 :     "A Different Exit Road",
            }
            City2 := City{
                path1 :     "A City Road",
                exit1 :     "A Exit Road",
                cityName :  "Napier",
                path2 :     "Another City Road",
                exit2 :     "A Different Exit Road",
            }

            if diff := deep.Equal(City1, City2); diff != nil {
                t.Error(diff)
            }
        }

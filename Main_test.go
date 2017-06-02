package main

import (
        "testing"
        "github.com/stretchr/testify/assert"
        //"github.com/go-test/deep"
        "reflect"
        //"fmt"
        )

//-------------------------------- Created Structs ---------------

func TestCityType(t *testing.T) {
// Set up

    City1 := City{"A City Road", "A Exit Road","Napier", "Another City Road","A Different Exit Road",}
    City2 := City{"A City Road", "A Exit Road","Napier", "Another City Road","A Different Exit Road",}

//Test
     if City1 != City2 {
       t.Error("The structures are not equal")
     }

    // if diff := deep.Equal(City1, City2); diff != nil {
    //   t.Error(diff)
    // }

}

func TestDriverType(t *testing.T){

//Set up
  var totalOutput []string
  City1 := City{"A City Road", "A Exit Road","Napier", "Another City Road","A Different Exit Road",}
  driver1 := Driver { 1, false, 6, 3, 1, City1, "Napier", "No Path", "No Exit", "Akina", totalOutput}
  driver2 := Driver { 1, false, 6, 3, 1, City1, "Napier", "No Path", "No Exit", "Akina", totalOutput}
// Test

  if  reflect.DeepEqual(driver1 , driver2) == false {
    t.Error("The structures are not equal")
  }
}

//--------------------------------akinaCount()----------------

func TestAkinaCountPlus(t *testing.T){
// Set Up

  var totalOutput []string
  City1 := City{"Blank","Blank","Blank","Blank","Blank",}
  driver := Driver { 1, false, 6, 3, 1, City1, "Akina", "", "", "", totalOutput}

  count := driver.akinaCount
  driver = akinaCount(driver)

// Test

  if driver.akinaCount != count + 1{
    t.Error("The count did not increase by 1")
  }
}


func TestAkinaCountStay(t *testing.T){
// Set Up

  var totalOutput []string
  City1 := City{"Blank","Blank","Blank","Blank","Blank",}
  driver := Driver { 1, false, 6, 3, 1, City1, "", "", "", "", totalOutput}

  count := driver.akinaCount
  driver = akinaCount(driver)

// Test

  if driver.akinaCount != count{
    t.Error("The count has changed")
  }
}


// --------------------------------driveExit()---------------------

func TestDriverExits(t *testing.T){
// Set up
  var totalOutput []string
  City1 := City{"Blank","Blank","Blank","Blank","Blank",}
  driver := Driver { 1, false, 2, 2, 1, City1, "Akina", "", "", "", totalOutput}

  driver = driveExit(driver)
// Test

  if driver.key != true {
    t.Error("The driver did not exit caused by the key")
    }
    if driver.nextCityName != "Outside City"{
      t.Error("The driver did not exit caused by the City Name")
    }
}

func TestDriverExitsFails(t *testing.T){
// Set up

  City1 := City{"Blank","Blank","Blank","Blank","Blank",}
  var totalOutput []string
  driver := Driver { 1, false, 6, 3, 1, City1, "Akina", "", "", "", totalOutput}

  driver = driveExit(driver)
// Test

  if driver.key != false {
    t.Error("The driver did exit. Casued by the key")
  }
  if driver.nextCityName == "Outside City"{
    t.Error("The driver did exit. Casued by the City Name")
  }
}

// --------------------------------driveCity()---------------------

func TestCityChangeUp(t *testing.T){
// Setup
  assert := assert.New(t)
  City1 := City {"Path1 City1","Exit1 City1","City1","Path2 City1","Exit2  City1"}
  City2 := City {"Path1 City2","Exit1 City2","City2","Path2 City2","Exit2  City2"}
  City3 := City {"Path1 City3","Exit1 City3","City3","Path2 City3","Exit2  City3"}
  City4 := City {"Path1 City4","Exit1 City4","City4","Path2 City4","Exit2  City4"}
  var allCities [4] City
  allCities[0] = City1
  allCities[1] = City2
  allCities[2] = City3
  allCities[3] = City4
  var totalOutput []string
  driver := Driver { 1, false, 1, 0, 0, City1, "", "", "", "", totalOutput}

  citynumber := driver.currentCityNumber
  driver = driveCity(driver , allCities)
// Test

  assert.Equal( citynumber + 1, driver.currentCityNumber , "The driver did not move to the City 1 above")

}

func TestCityChangeDown(t *testing.T){
  // Setup
    assert := assert.New(t)
    City1 := City {"Path1 City1","Exit1 City1","City1","Path2 City1","Exit2  City1"}
    City2 := City {"Path1 City2","Exit1 City2","City2","Path2 City2","Exit2  City2"}
    City3 := City {"Path1 City3","Exit1 City3","City3","Path2 City3","Exit2  City3"}
    City4 := City {"Path1 City4","Exit1 City4","City4","Path2 City4","Exit2  City4"}
    var allCities [4] City
    allCities[0] = City1
    allCities[1] = City2
    allCities[2] = City3
    allCities[3] = City4
    var totalOutput []string
    driver := Driver { 1, false, 0, 0, 1, City2, "", "", "", "", totalOutput}

    citynumber := driver.currentCityNumber
    driver = driveCity(driver , allCities)
  // Test

    assert.Equal(citynumber - 1 , driver.currentCityNumber , "The driver did not move to the City 1 below")

}

func TestCitySkipUp(t *testing.T){ // Actually tests for the number being set lower
  // Setup
    assert := assert.New(t)
    City1 := City {"Path1 City1","Exit1 City1","City1","Path2 City1","Exit2  City1"}
    City2 := City {"Path1 City2","Exit1 City2","City2","Path2 City2","Exit2  City2"}
    City3 := City {"Path1 City3","Exit1 City3","City3","Path2 City3","Exit2  City3"}
    City4 := City {"Path1 City4","Exit1 City4","City4","Path2 City4","Exit2  City4"}
    var allCities [4] City
    allCities[0] = City1
    allCities[1] = City2
    allCities[2] = City3
    allCities[3] = City4
    var totalOutput []string
    driver := Driver { 1, false, 1, 0, 3, City3, "", "", "", "", totalOutput}

    driver = driveCity(driver , allCities)
  // Test

    assert.Equal( 0 ,driver.currentCityNumber, "The driver did not skip Cities ")

}
func TestCitySkipDown(t *testing.T){ // Actually tests for the number being set higher
  // Setup
    assert := assert.New(t)
    City1 := City {"Path1 City1","Exit1 City1","City1","Path2 City1","Exit2  City1"}
    City2 := City {"Path1 City2","Exit1 City2","City2","Path2 City2","Exit2  City2"}
    City3 := City {"Path1 City3","Exit1 City3","City3","Path2 City3","Exit2  City3"}
    City4 := City {"Path1 City4","Exit1 City4","City4","Path2 City4","Exit2  City4"}
    var allCities [4] City
    allCities[0] = City1
    allCities[1] = City2
    allCities[2] = City3
    allCities[3] = City4
    var totalOutput []string
    driver := Driver { 1, false, 0, 0, 0, City1, "", "", "", "", totalOutput}

    driver = driveCity(driver , allCities)
  // Test

    assert.Equal( 3, driver.currentCityNumber,"The driver did not skip Cities ")

}

//--------------------------------driveCityPrint()-----------------

func TestAddingToPrintMain(t *testing.T){
//Set Up
  assert := assert.New(t)
  City2 := City{"Blank","Blank","Blank","Blank","Blank",}
  var totalOutput []string
  driver := Driver { 1, false, 0, 2, 1, City2, "Akina", "the Road", "-", "NextCity", totalOutput}
  driver = driveCityPrint(driver)
// Test

  assert.Equal("Driver 1 heading from Akina to NextCity via the Road", driver.totalOutput[0],"Values were not the same")

}

func TestAddingToPrintExtra(t *testing.T){
//Set Up
  assert := assert.New(t)
  City2 := City{"Blank","Blank","Blank","Blank","Blank",}
  var totalOutput []string
  driver := Driver { 1, false, 0, 2, 1, City2, "Akina", "-", "Karamu Road", "Outside City", totalOutput}
  driver = driveCityPrint(driver)
// Test

  assert.Equal("Driver 1 has gone to Napier", driver.totalOutput[1],"Values were not the same")

}

//--------------------------------extras()-------------------------

func TestMeetingExtra(t *testing.T){
//Set Up
  assert := assert.New(t)
  City2 := City{"Blank","Blank","Blank","Blank","Blank",}
  var totalOutput []string
  driver := Driver { 1, false, 0, 5, 0, City2, "-", "-", "-", "-", totalOutput}
  driver = extras(driver)
// Test

    assert.Equal("Driver 1 Met With Jhon Jamieson 5 Times!", driver.totalOutput[0],"Values were not the same")

}

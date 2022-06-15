package microkit

import (
	"testing"
	"reflect"
	"fmt"
)

func TestMicroKit(t *testing.T)  {
	jsonv := `{"config":"dc5+idVVGi00GpKlmkp5PyMHZ0AzpZuyl+0sfTVzYcFHCc\/8bMGvamSA3jfYEfJZJ6i6eMu4vHIvpqBF+\/LH4m9tO6Armcm4vSk9cM1rlrtypLA\/q+Ms1EUgXBWdoeJbZOz\/k8QvRu7VGz\/s9RtFPL3ouzMdKaFAj78hIMscQG5uv1ClZoLNXbLrAzO9sJPmyFS4GfUulgZ2\/sih3G7qNVaMclHcIVX6vtvz09pEF7CRH1SMZCXd\/ceh\/e1ifpeZk1g762TNF37NBIDFMf0f50dvpM0OAo+U8+WZtJASPriD4zPVvp5mGR+LdOrPYFFHpby+cyAqutWMo16a1qTIkPcQkt0XRcIGlANrCXjiCwZ8c+Vu2z3feLPurq1E1HTiTwPVB5zRuaVYstLqDmEkV3AZZI0IMAL\/nHKJFLGkFeWfVQNBlCwgeqrbSqLDh5Pn3m6\/N3MEiypSivKGjRsUSMsyIV2nrmPVUMvE3BMkXsTA0u0RYSMGCIQhqf1RbtypCfrnPIsJVA7+JSHfkMx3ZIA+1BPL9n+Hbvb4RQxPwXp+3XXkhCzKfdc2loWxMFCV\/NSi9XcGrFATLAYb9tz4\/aNhzXU+WWOLa3+RflLLSDHFKvz18pTOCNdmK4fujLv5mgH0KmNpWnndFRWY\/mJjuetBxov8NolSFrIXUcE\/MzqEc2IeB68Fv9AXMw1Uz62+Cq3rQVeTnTlattRNrtamWBluUMqA\/bZKIUP1LV\/gF8kSVSVzLWZ\/iHg1mr1YxYswkoVoKIpp3qBmkN6CyKtyHNmyRSxL\/d5fUfi5tUHmbZ\/1lMMtA0v3KC2up+uC245uup1V9H++6YHe392FaYyZe41C5A3V3vFMJiTC2VnSd\/ZbmjNLfwhPvUqcfwLApYwyXWjwPd2gNNM2lL+VVhsjTuCFNGZ+YccSKQXW4EJzUjg2YB2R\/MqFEYePL\/OyAv4K"}`
	uri := "3"
	port := 8083
	MockServer(jsonv, port, uri)
	microKit := InitKit("3c20444b-6865-4dfb-bfee-0a0b7279572b-qPYBtP5lgv4f/TDOV17BZA==", map[string]string{"url": fmt.Sprintf("http://localhost:%d/%s", port, uri)})
	if reflect.TypeOf(*microKit).String() != "microkit.MicroKit" {
		t.Errorf("got %s want %s", reflect.TypeOf(*microKit), "microkit.MicroKit")
	}
}

func TestMicroKitReady(t *testing.T)  {
	jsonv := `{"config":"dc5+idVVGi00GpKlmkp5PyMHZ0AzpZuyl+0sfTVzYcFHCc\/8bMGvamSA3jfYEfJZJ6i6eMu4vHIvpqBF+\/LH4m9tO6Armcm4vSk9cM1rlrtypLA\/q+Ms1EUgXBWdoeJbZOz\/k8QvRu7VGz\/s9RtFPL3ouzMdKaFAj78hIMscQG5uv1ClZoLNXbLrAzO9sJPmyFS4GfUulgZ2\/sih3G7qNVaMclHcIVX6vtvz09pEF7CRH1SMZCXd\/ceh\/e1ifpeZk1g762TNF37NBIDFMf0f50dvpM0OAo+U8+WZtJASPriD4zPVvp5mGR+LdOrPYFFHpby+cyAqutWMo16a1qTIkPcQkt0XRcIGlANrCXjiCwZ8c+Vu2z3feLPurq1E1HTiTwPVB5zRuaVYstLqDmEkV3AZZI0IMAL\/nHKJFLGkFeWfVQNBlCwgeqrbSqLDh5Pn3m6\/N3MEiypSivKGjRsUSMsyIV2nrmPVUMvE3BMkXsTA0u0RYSMGCIQhqf1RbtypCfrnPIsJVA7+JSHfkMx3ZIA+1BPL9n+Hbvb4RQxPwXp+3XXkhCzKfdc2loWxMFCV\/NSi9XcGrFATLAYb9tz4\/aNhzXU+WWOLa3+RflLLSDHFKvz18pTOCNdmK4fujLv5mgH0KmNpWnndFRWY\/mJjuetBxov8NolSFrIXUcE\/MzqEc2IeB68Fv9AXMw1Uz62+Cq3rQVeTnTlattRNrtamWBluUMqA\/bZKIUP1LV\/gF8kSVSVzLWZ\/iHg1mr1YxYswkoVoKIpp3qBmkN6CyKtyHNmyRSxL\/d5fUfi5tUHmbZ\/1lMMtA0v3KC2up+uC245uup1V9H++6YHe392FaYyZe41C5A3V3vFMJiTC2VnSd\/ZbmjNLfwhPvUqcfwLApYwyXWjwPd2gNNM2lL+VVhsjTuCFNGZ+YccSKQXW4EJzUjg2YB2R\/MqFEYePL\/OyAv4K"}`
	uri := "4"
	port := 8084
	MockServer(jsonv, port, uri)
	microKit := InitKit("3c20444b-6865-4dfb-bfee-0a0b7279572b-qPYBtP5lgv4f/TDOV17BZA==", map[string]string{"url": fmt.Sprintf("http://localhost:%d/%s", port, uri)})
	
	if !microKit.KitReady  {
		t.Errorf("got %t want %t", microKit.KitReady, true)
	}
}

func TestMicroKitConfigKitType(t *testing.T)  {
	jsonv := `{"config":"dc5+idVVGi00GpKlmkp5PyMHZ0AzpZuyl+0sfTVzYcFHCc\/8bMGvamSA3jfYEfJZJ6i6eMu4vHIvpqBF+\/LH4m9tO6Armcm4vSk9cM1rlrtypLA\/q+Ms1EUgXBWdoeJbZOz\/k8QvRu7VGz\/s9RtFPL3ouzMdKaFAj78hIMscQG5uv1ClZoLNXbLrAzO9sJPmyFS4GfUulgZ2\/sih3G7qNVaMclHcIVX6vtvz09pEF7CRH1SMZCXd\/ceh\/e1ifpeZk1g762TNF37NBIDFMf0f50dvpM0OAo+U8+WZtJASPriD4zPVvp5mGR+LdOrPYFFHpby+cyAqutWMo16a1qTIkPcQkt0XRcIGlANrCXjiCwZ8c+Vu2z3feLPurq1E1HTiTwPVB5zRuaVYstLqDmEkV3AZZI0IMAL\/nHKJFLGkFeWfVQNBlCwgeqrbSqLDh5Pn3m6\/N3MEiypSivKGjRsUSMsyIV2nrmPVUMvE3BMkXsTA0u0RYSMGCIQhqf1RbtypCfrnPIsJVA7+JSHfkMx3ZIA+1BPL9n+Hbvb4RQxPwXp+3XXkhCzKfdc2loWxMFCV\/NSi9XcGrFATLAYb9tz4\/aNhzXU+WWOLa3+RflLLSDHFKvz18pTOCNdmK4fujLv5mgH0KmNpWnndFRWY\/mJjuetBxov8NolSFrIXUcE\/MzqEc2IeB68Fv9AXMw1Uz62+Cq3rQVeTnTlattRNrtamWBluUMqA\/bZKIUP1LV\/gF8kSVSVzLWZ\/iHg1mr1YxYswkoVoKIpp3qBmkN6CyKtyHNmyRSxL\/d5fUfi5tUHmbZ\/1lMMtA0v3KC2up+uC245uup1V9H++6YHe392FaYyZe41C5A3V3vFMJiTC2VnSd\/ZbmjNLfwhPvUqcfwLApYwyXWjwPd2gNNM2lL+VVhsjTuCFNGZ+YccSKQXW4EJzUjg2YB2R\/MqFEYePL\/OyAv4K"}`
	uri := "5"
	port := 8085
	MockServer(jsonv, port, uri)
	microKit := InitKit("3c20444b-6865-4dfb-bfee-0a0b7279572b-qPYBtP5lgv4f/TDOV17BZA==", map[string]string{"url": fmt.Sprintf("http://localhost:%d/%s", port, uri)})
	
	if reflect.TypeOf(microKit.ConfigKit).String() != "map[string]microkit.ConfigGroup" {
		t.Errorf("got %s want %s", reflect.TypeOf(microKit.ConfigKit), "map[string]microkit.ConfigGroup")
	}
}

func TestMicroKitConfigKitValue(t *testing.T)  {
	jsonv := `{"config":"dc5+idVVGi00GpKlmkp5PyMHZ0AzpZuyl+0sfTVzYcFHCc\/8bMGvamSA3jfYEfJZJ6i6eMu4vHIvpqBF+\/LH4m9tO6Armcm4vSk9cM1rlrtypLA\/q+Ms1EUgXBWdoeJbZOz\/k8QvRu7VGz\/s9RtFPL3ouzMdKaFAj78hIMscQG5uv1ClZoLNXbLrAzO9sJPmyFS4GfUulgZ2\/sih3G7qNVaMclHcIVX6vtvz09pEF7CRH1SMZCXd\/ceh\/e1ifpeZk1g762TNF37NBIDFMf0f50dvpM0OAo+U8+WZtJASPriD4zPVvp5mGR+LdOrPYFFHpby+cyAqutWMo16a1qTIkPcQkt0XRcIGlANrCXjiCwZ8c+Vu2z3feLPurq1E1HTiTwPVB5zRuaVYstLqDmEkV3AZZI0IMAL\/nHKJFLGkFeWfVQNBlCwgeqrbSqLDh5Pn3m6\/N3MEiypSivKGjRsUSMsyIV2nrmPVUMvE3BMkXsTA0u0RYSMGCIQhqf1RbtypCfrnPIsJVA7+JSHfkMx3ZIA+1BPL9n+Hbvb4RQxPwXp+3XXkhCzKfdc2loWxMFCV\/NSi9XcGrFATLAYb9tz4\/aNhzXU+WWOLa3+RflLLSDHFKvz18pTOCNdmK4fujLv5mgH0KmNpWnndFRWY\/mJjuetBxov8NolSFrIXUcE\/MzqEc2IeB68Fv9AXMw1Uz62+Cq3rQVeTnTlattRNrtamWBluUMqA\/bZKIUP1LV\/gF8kSVSVzLWZ\/iHg1mr1YxYswkoVoKIpp3qBmkN6CyKtyHNmyRSxL\/d5fUfi5tUHmbZ\/1lMMtA0v3KC2up+uC245uup1V9H++6YHe392FaYyZe41C5A3V3vFMJiTC2VnSd\/ZbmjNLfwhPvUqcfwLApYwyXWjwPd2gNNM2lL+VVhsjTuCFNGZ+YccSKQXW4EJzUjg2YB2R\/MqFEYePL\/OyAv4K"}`
	uri := "6"
	port := 8086
	MockServer(jsonv, port, uri)
	microKit := InitKit("3c20444b-6865-4dfb-bfee-0a0b7279572b-qPYBtP5lgv4f/TDOV17BZA==", map[string]string{"url": fmt.Sprintf("http://localhost:%d/%s", port, uri)})
	
	if reflect.TypeOf(microKit.ConfigKit["db"]).String() != "microkit.ConfigGroup" {
		t.Errorf("got %s want %s", reflect.TypeOf(microKit.ConfigKit["db"]), "microkit.ConfigGroup")
	}
}
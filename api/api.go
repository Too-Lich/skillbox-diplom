package api

import (
	"diplom/internal/storages"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {
		//r.Get("/", ConnectionHandler)

	})
	r.Get("/", ConnectionHandler)
	r.Get("/test", handleTest)
	return r
}

func ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	res := storages.GetResultData()
	resBytes, _ := json.Marshal(res)
	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resBytes)
	//log.Print(string(resBytes))
	//w.Write([]byte("{\n  \"status\": true,\n  \"data\": {},\n  \"error\": \"\"\n}"))

}

func handleTest(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\n  \"status\": true,\n  \"data\": {\n    \"sms\": [\n      [\n        {\n          \"country\": \"Cuba\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ],\n      [\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Canada\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ]\n    ],\n    \"mms\": [\n      [\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Canada\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ],\n      [\n        {\n          \"country\": \"Canada\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ]\n    ],\n    \"voice_call\": [\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"TransparentCalls\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      },\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"TransparentCalls\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      },\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"E-Voice\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      },\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"E-Voice\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      }\n    ],\n    \"email\": [\n      [\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 195\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        }\n      ],\n      [\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        }\n      ]\n    ],\n    \"billing\": {\n      \"create_customer\": true,\n      \"purchase\": true,\n      \"payout\": true,\n      \"recurring\": false,\n      \"fraud_control\": true,\n      \"checkout_page\": false\n    },\n    \"support\": [\n      3,\n      62\n    ],\n    \"incident\": [\n      {\"topic\":  \"Topic 1\", \"status\": \"active\"},\n      {\"topic\":  \"Topic 2\", \"status\": \"active\"},\n      {\"topic\":  \"Topic 3\", \"status\": \"closed\"},\n      {\"topic\":  \"Topic 4\", \"status\": \"closed\"}\n    ]\n  },\n  \"error\": \"\"\n}"))
}

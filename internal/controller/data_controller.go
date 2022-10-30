package controller

import (
	"log"
	"path/filepath"
	"techTask/config"
	"techTask/internal/models"
	"techTask/internal/services"
)

type DataController struct {
	services *services.Manager
	cfg      *config.Config
}

func NewDataController(services *services.Manager) *DataController {
	return &DataController{
		services: services,
	}
}

func (ctr *DataController) DownloadFile(url, filename string) error {

	err := ctr.services.Data.DownloadFileService(url, filename)

	if err != nil {
		return err
	}
	path, _ := filepath.Abs(filename)
	err = ctr.services.Data.StartInsertService(path)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (ctr *DataController) Search(q *models.Queries) (*[]models.DataParse, error) {
	data, err := ctr.services.Data.SearchService(q)
	return data, err
}

//func (ctr *DataController) Welcome(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode("Welcome")
//	default:
//		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
//	}
//}
//func (ctr *DataController) GetByTransactionID(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//
//		id, err := strconv.Atoi(chi.URLParam(r, "id"))
//		if err != nil {
//			log.Println(err)
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//
//		data, err := ctr.services.Data.GetByTransactionIDService(id)
//		if data == nil {
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//		d, err := json.Marshal(data)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//
//		fmt.Println(string(d))
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(data)
//		return
//
//	default:
//		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
//	}
//
//}
//
//func (ctr *DataController) GetByTerminalID(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//		id, err := strconv.Atoi(chi.URLParam(r, "id"))
//		if err != nil {
//			log.Println(err)
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		data, err := ctr.services.Data.GetByTerminalIDService(id)
//		fmt.Println(err, data)
//		if data == nil {
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//
//		d, err := json.Marshal(data)
//		if err != nil {
//			log.Println(err)
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//
//		fmt.Println(string(d))
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(data)
//		return
//
//	default:
//		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
//	}
//
//}
//
//func (ctr *DataController) GetByStatus(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//		status := chi.URLParam(r, "s")
//
//		data, err := ctr.services.Data.GetByStatusService(status)
//		if data == nil {
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//
//		d, err := json.Marshal(data)
//		if err != nil {
//			log.Println(err)
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//
//		fmt.Println(string(d))
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(data)
//		return
//
//	default:
//		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
//	}
//
//}
//
//func (ctr *DataController) GetByPaymentType(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//		payType := chi.URLParam(r, "t")
//
//		data, err := ctr.services.Data.GetByPaymentTypeService(payType)
//		if data == nil {
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//		fmt.Println(payType)
//		if err != nil {
//			log.Println(err)
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		d, err := json.Marshal(data)
//		if err != nil {
//			log.Println(err)
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//
//		fmt.Println(string(d))
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(data)
//		return
//
//	default:
//		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
//	}
//
//}
//
//func (ctr *DataController) GetByTime(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//		start := chi.URLParam(r, "start")
//		end := chi.URLParam(r, "end")
//
//		fmt.Println(start, end)
//		s, err := time.Parse("2006-01-02", start)
//		if err != nil {
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//		e, err := time.Parse("2006-01-02", end)
//		if err != nil {
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//		si := s.Unix()
//		ei := e.Unix()
//
//		if si > ei {
//			si, ei = ei, si
//		}
//		fmt.Println(start, end)
//		data, err := ctr.services.Data.GetByDataPostService(si, ei)
//		if data == nil {
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//		if err != nil {
//			log.Println(err)
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		d, err := json.Marshal(data)
//		if err != nil {
//			log.Println(err)
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//
//		fmt.Println(string(d))
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(data)
//		return
//
//	default:
//		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
//	}
//
//}
//func (ctr *DataController) GetByNarrative(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//		nar := chi.URLParam(r, "text")
//		fmt.Println(nar)
//		data, err := ctr.services.Data.GetByPaymentNarrativeService(nar)
//
//		fmt.Println(data)
//		if data == nil {
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//
//		if err != nil {
//			log.Println(err)
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		d, err := json.Marshal(data)
//		if err != nil {
//			log.Println(err)
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//
//		fmt.Println(string(d))
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(data)
//		return
//
//	default:
//		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
//	}
//
//}

//
//func (ctr *DataController) Search(c *gin.Context) {
//	switch r.Method {
//	case http.MethodGet:
//		fmt.Println(123)
//		fmt.Println(chi.RouteContext(r.Context()).RoutePatterns)
//		p := chi.URLParam(r, "search")
//		fmt.Println(p)
//
//	default:
//		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
//
//	}
//}

package vm_stat

type MemoryStats struct {
	PagesFree                 string `json:"pages_free"`
	PagesActive               string `json:"pages_active"`
	PagesInactive             string `json:"pages_inactive"`
	PagesSpeculative          string `json:"pages_speculative"`
	PagesThrottled            string `json:"pages_throttled"`
	PagesWiredDown            string `json:"pages_wired_down"`
	PagesPurgeable            string `json:"pages_purgeable"`
	TranslationFaults         string `json:"translation_faults"`
	PagesCopyOnWrite          string `json:"pages_copy_on_write"`
	PagesZeroFilled           string `json:"pages_zero_filled"`
	PagesReactivated          string `json:"pages_reactivated"`
	PagesPurged               string `json:"pages_purged"`
	FileBackedPages           string `json:"file_backed_pages"`
	AnonymousPages            string `json:"anonymous_pages"`
	PagesStoredInCompressor   string `json:"pages_stored_in_compressor"`
	PagesOccupiedByCompressor string `json:"pages_occupied_by_compressor"`
	Decompressions            string `json:"decompressions"`
	Compressions              string `json:"compressions"`
	Pageins                   string `json:"pageins"`
	Pageouts                  string `json:"pageouts"`
	Swapins                   string `json:"swapins"`
	Swapouts                  string `json:"swapouts"`
}

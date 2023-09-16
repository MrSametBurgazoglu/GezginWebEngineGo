graph:
	godepgraph -novendor -s -p github.com,golang.org,gorm.io,gopkg ./main/ | dot -Tpng -o godepgraph.png
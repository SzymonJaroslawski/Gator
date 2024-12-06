module github.com/SzymonJaroslawski/Gator

go 1.23.3

require github.com/SzymonJaroslawski/Gator/internal/config v0.0.0

require github.com/SzymonJaroslawski/Gator/internal/database v0.0.0

require github.com/SzymonJaroslawski/Gator/internal/RSS v0.0.0

require (
	github.com/golang/glog v1.2.3
	github.com/lib/pq v1.10.9
)

require github.com/google/uuid v1.6.0 // indirect

replace github.com/SzymonJaroslawski/Gator/internal/config v0.0.0 => ./internal/config/

replace github.com/SzymonJaroslawski/Gator/internal/database v0.0.0 => ./internal/database/

replace github.com/SzymonJaroslawski/Gator/internal/RSS v0.0.0 => ./internal/RSS/

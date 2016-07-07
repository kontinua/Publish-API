package cmd

import (
	"net"

	flag "github.com/spf13/pflag"

	"k8s.io/minikube/pkg/localkube"
	"k8s.io/minikube/pkg/util"
)

func NewLocalkubeServer() *localkube.LocalkubeServer {
	// net.ParseCIDR returns multiple values. Use the IPNet return value
	_, defaultServiceClusterIPRange, _ := net.ParseCIDR(util.DefaultServiceClusterIP + "/24")

	return &localkube.LocalkubeServer{
		Containerized:            false,
		EnableDNS:                true,
		DNSDomain:                util.DefaultDNSDomain,
		DNSIP:                    net.ParseIP(util.DefaultDNSIP),
		LocalkubeDirectory:       util.DefaultLocalkubeDirectory,
		ServiceClusterIPRange:    *defaultServiceClusterIPRange,
		APIServerAddress:         net.ParseIP("0.0.0.0"),
		APIServerPort:            443,
		APIServerInsecureAddress: net.ParseIP("127.0.0.1"),
		APIServerInsecurePort:    8080,
		ShouldGenerateCerts:      true,
		ShowVersion:              false,
	}
}

// AddFlags adds flags for a specific LocalkubeServer
func AddFlags(s *localkube.LocalkubeServer) {
	flag.BoolVar(&s.Containerized, "containerized", s.Containerized, "If kubelet should run in containerized mode")
	flag.BoolVar(&s.EnableDNS, "enable-dns", s.EnableDNS, "If dns should be enabled")
	flag.StringVar(&s.DNSDomain, "dns-domain", s.DNSDomain, "The cluster dns domain")
	flag.IPVar(&s.DNSIP, "dns-ip", s.DNSIP, "The cluster dns IP")
	flag.StringVar(&s.LocalkubeDirectory, "localkube-directory", s.LocalkubeDirectory, "The directory localkube will store files in")
	flag.IPNetVar(&s.ServiceClusterIPRange, "service-cluster-ip-range", s.ServiceClusterIPRange, "The service-cluster-ip-range for the apiserver")
	flag.IPVar(&s.APIServerAddress, "apiserver-address", s.APIServerAddress, "The address the apiserver will listen securely on")
	flag.IntVar(&s.APIServerPort, "apiserver-port", s.APIServerPort, "The port the apiserver will listen securely on")
	flag.IPVar(&s.APIServerInsecureAddress, "apiserver-insecure-address", s.APIServerInsecureAddress, "The address the apiserver will listen insecurely on")
	flag.IntVar(&s.APIServerInsecurePort, "apiserver-insecure-port", s.APIServerInsecurePort, "The port the apiserver will listen insecurely on")
	flag.BoolVar(&s.ShouldGenerateCerts, "generate-certs", s.ShouldGenerateCerts, "If localkube should generate it's own certificates")
	flag.BoolVar(&s.ShowVersion, "version", s.ShowVersion, "If localkube should just print the version and exit.")

	// These two come from vendor/ packages that use flags. We should hide them
	flag.CommandLine.MarkHidden("google-json-key")
	flag.CommandLine.MarkHidden("log-flush-frequency")

	// Parse them
	flag.Parse()
}

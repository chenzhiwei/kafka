package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type Config struct {
	Broker         string
	SkipTLSVerify  bool
	CaFile         string
	ClientKeyFile  string
	ClientCertFile string
	Mechanism      string
	Username       string
	Password       string
}

func (c *Config) dialer() (*kafka.Dialer, error) {
	config := &tls.Config{}

	if c.SkipTLSVerify {
		config.InsecureSkipVerify = true
	}

	if c.CaFile != "" {
		caCert, err := os.ReadFile(c.CaFile)
		if err != nil {
			return nil, err
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		config.RootCAs = caCertPool
	}

	if c.ClientCertFile != "" && c.ClientKeyFile != "" {
		clientCertificate, err := tls.LoadX509KeyPair(c.ClientCertFile, c.ClientKeyFile)
		if err != nil {
			return nil, err
		}

		config.Certificates = []tls.Certificate{clientCertificate}
	}

	mechanism, err := c.mechanism()
	if err != nil {
		return nil, err
	}

	return &kafka.Dialer{
		DualStack:     true,
		TLS:           config,
		SASLMechanism: mechanism,
	}, nil
}

func (c *Config) mechanism() (sasl.Mechanism, error) {
	switch c.Mechanism {
	case "PLAIN":
		return plain.Mechanism{
			Username: c.Username,
			Password: c.Password,
		}, nil
	case "SCRAM-SHA-256":
		return scram.Mechanism(scram.SHA256, c.Username, c.Password)
	case "SCRAM-SHA-512":
		return scram.Mechanism(scram.SHA256, c.Username, c.Password)
	default:
		return nil, nil
	}
}

func (c *Config) Connection() (*kafka.Conn, error) {
	dialer, err := c.dialer()
	if err != nil {
		return nil, err
	}

	var conn *kafka.Conn

	brokers := strings.Split(c.Broker, ",")
	for _, addr := range brokers {
		if conn, err = dialer.Dial("tcp", addr); err != nil {
			continue
		} else {
			break
		}
	}

	if conn == nil {
		fmt.Printf("unable to dial broker: %s, error: %v\n", c.Broker, err)
		return nil, err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return nil, err
	}

	ctrConn, err := dialer.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		fmt.Printf("unable to dial controller: %v\n", err)
		return nil, err
	}
	defer ctrConn.Close()

	return ctrConn, nil
}

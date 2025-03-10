# Proto Package

This repository contains a Go package for handling the encoding and decoding of protocol message frames. It includes message construction, payload handling, and register manipulation. The package implements features such as message construction, payload encoding, checksum calculation, and decoding from raw frame data.

It goes with the side project [github.com/kumy/samsung-ac-simulator](https://github.com/kumy/samsung-ac-simulator)


## Features

- **Message Construction**: Create messages with counters, payloads, checksums, and register details.
- **Payload Management**: Manipulate payloads with registers and various types of data encoding.
- **Hex Encoding**: Supports converting messages and payloads to hexadecimal format.
- **Register Handling**: Register management with the ability to add, remove, and manipulate registers.
- **Checksum Validation**: Supports XOR-based checksum calculation for verifying message integrity.
- **Decoding**: Decode raw frame data into structured message objects, extracting payloads and registers.

## Installation

To install and use the package, ensure you have Go installed on your machine. Then, run the following command:

```bash
go get github.com/yourusername/proto
```

## Usage

### Encoding a Message

To create and encode a message, you can use the `message` struct and its methods:

```go
package main

import (
	"fmt"
	"github.com/kumy/samsung-ac/proto"
)

func main() {
	// Create a new counter and message
	counter := proto.NewCounter(1)
	msg := proto.NewEmptyMessage(counter)

	// Set the payload and counter
	payload := proto.NewPayload(proto.MessageType{0x1202})
	msg.SetPayload(payload)

	// Encode the message to bytes
	encodedMsg := msg.Bytes()

	// Print the encoded message
	fmt.Printf("Encoded Message: %x\n", encodedMsg)
}
```

### Decoding a Message

To decode a message from a raw byte frame, use the `decoder`:

```go
package main

import (
	"fmt"
	"github.com/kumy/samsung-ac/proto"
)

func main() {
	// Example raw frame (hexadecimal string)
	frame := "your-hex-frame-data-here"

	// Decode the message
	decoder := proto.NewDecoder()
	msg := decoder.DecodeFromString(frame)

	// Print the decoded message
	fmt.Printf("Decoded Message: %s\n", msg.String())
}
```

### Register Handling

You can manipulate registers using the `Register` interface:

```go
package main

import (
	"fmt"
	"github.com/kumy/samsung-ac/proto"
)

func main() {
	// Create a new register
	register := proto.NewRegister(0x01, []byte{0x01, 0x02})

	// Print register details
	fmt.Printf("Register: %s\n", register.String())
}
```

## Contributing

Contributions are welcome! If you'd like to contribute to this project, feel free to fork the repository, create a branch, and submit a pull request.

Please ensure that your changes are well-tested and follow the existing code style.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
know if you need any modifications or additional information in the `README.md`.
# zfetch

zfetch is a lightweight system information tool and a cool alternative to neofetch, written in Go. With zfetch, you can quickly glance at your system information and showcase it in style.

![image](https://github.com/zubairmh/zfetch/assets/116816535/d6ff288c-450f-4692-ab5e-40532c69d416)
## Features

- System Information: zfetch provides a comprehensive overview of your system, including the hostname, CPU, memory usage, disk usage and uptime
- Lightweight: zfetch is designed to be fast and lightweight, ensuring it doesn't impact your system's performance.
- Cross-Platform: zfetch works on major operating systems, including Linux, macOS, and Windows.

## Installation

To use zfetch, ensure you have Go installed on your system. Then, follow the steps below:

Clone the zfetch repository to your local machine.

    git clone https://github.com/zubairmh/zfetch.git
    cd zfetch
    go mod tidy

Build the zfetch binary and add it to path

    go build
    mv ./zfetch ~/.local/bin

## Usage

To run zfetch, open your terminal and execute the following command:

`zfetch`

## Contributions 

If you'd like to contribute support for additional operating systems or improve the existing ones, feel free to submit a pull request.

## License

This project is licensed under the [MIT License](https://github.com/zubairmh/zfetch/blob/main/LICENSE).

## Acknowledgements

Without them, this project wouldn't be possible

- The Go programming language community for providing a powerful and efficient language for building zfetch.
- The [psutil](https://github.com/shirou/gopsutil) library, for providing an extremely useful API to interact with the system

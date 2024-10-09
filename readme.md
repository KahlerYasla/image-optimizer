# image-optimizer

> A simple image optimizer tool to convert images to WebP format using Go.

This Go application converts images from JPEG, PNG, and WebP formats to WebP format. The tool is designed for batch processing images within a specified directory.

## Table of Contents

-   [Requirements](#requirements)
-   [Installation](#installation)
-   [Usage](#usage)
-   [How It Works](#how-it-works)
-   [Functionality](#functionality)
-   [Error Handling](#error-handling)
-   [Contributing](#contributing)
-   [License](#license)

## Requirements

-   You can simply install the required packages by running the following command:

    ```bash
    go mod tidy
    ```

-   Go (version 1.14 or higher)
-   Required packages:
    -   [chai2010/webp](https://github.com/chai2010/webp)
    -   [golang.org/x/image](https://pkg.go.dev/golang.org/x/image/draw)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/KahlerYasla/image-optimizer.git
    cd image-optimizer
    ```

2. Install the required Go packages:

    ```bash
    go get -u github.com/chai2010/webp
    go get -u golang.org/x/image
    ```

3. Build the application:

    ```bash
    go build -o image-optimizer
    ```

## Usage

1. Run the application:

    ```bash
    ./image-optimizer
    ```

2. When prompted, enter the directory containing the images you want to convert.

3. The program will process each image in the directory, converting it to the WebP format.

## How It Works

-   The application scans the specified directory for images with the extensions `.jpeg`, `.jpg`, `.png`, or `.webp`.
-   Each valid image file is read and decoded using the appropriate image decoder.
-   The image is then optionally resized and converted to the WebP format, saved in the same directory with a `.webp` extension.

## Functionality

-   **Supported Formats:** JPEG, PNG, and WebP
-   **Output Format:** WebP
-   **Image Resizing:** The images are resized to maintain aspect ratio, optimizing for file size.

### Code Breakdown

-   **`main()` Function:**
    -   Prompts for a directory input.
    -   Reads the directory and processes each image file.
-   **`convertToWebP(inputPath, outputPath string) error`:**
    -   Handles image decoding, resizing, and encoding to WebP.

## Error Handling

The application includes basic error handling for:

-   Directory reading errors
-   Unsupported file formats
-   File I/O errors during reading and writing

Errors encountered during processing are logged to the console.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to customize this README further to suit your project's specific details or requirements!

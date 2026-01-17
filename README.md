# Web Wake-on-LAN

A simple web application to wake up PCs on your local network (Wake-on-LAN).
It runs as a single binary (Go backend) with an embedded React frontend.

## Usage

1.  **Download** the binary for your platform (or build it from source).

2.  **Run** the application with your target devices' names and MAC addresses as arguments:

    ```bash
    # Format: ./web-wol "Name=MAC_Address" ["Name=MAC_Address" ...]

    ./web-wol "MyPC=AA:BB:CC:DD:EE:FF" "HomeServer=11:22:33:44:55:66"
    ```

3.  **Open** your browser and go to:
    **`http://localhost:8080`**

4.  **Click** the "Wake" button to turn on your device.

---

## Build from Source

**Prerequisites:**
*   Go 1.20+
*   Node.js & pnpm

**Steps:**

1.  Clone the repository:
    ```bash
    git clone https://github.com/yhiraki/wakeonlan-webapp.git
    cd wakeonlan-webapp
    ```

2.  Build the frontend assets:
    ```bash
    cd frontend
    pnpm install
    pnpm build
    cd ..
    ```

3.  Build the single binary:
    ```bash
    # Build everything (Frontend + Backend)
    make build
    ```

## License

[MIT License](LICENSE)

# vectytemplater
Executable that generates a basic vecty application folder structure and files.

## Usage
1. Install `vectytemplater` by running

    ```shell
    go install github.com/soypat/vectytemplater@latest
    ```

2. Run `vectytemplater` in the parent directory where you want to create the folder with the first argument being name of folder.

    ```shell
    vectytemplater myproject
    ```
3. Run `wasmserve` in the newly created directory! Navigate your browser to [`localhost:8080`](http://localhost:8080/) to see the project. Install by running `go install github.com/hajimehoshi/wasmserve@latest`.
    ```shell
    wasmserve
    ```

4. Modify the project to your needs. Here are some initial suggestions:

    * Modify the module name by replacing the default value `vecty-templater-project` with your module's name in all project files.

    * Delete the `.vscode` folder if you are not using Visual Studio Code. It is there for intellisense to work since WASM projects targets the browser, not your local computer environment. WASM projects are compiled with Go environent variables `GOOS=js` and `GOARCH=wasm`.

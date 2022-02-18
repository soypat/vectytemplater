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

3. Modify the project to your needs. Here are some initial suggestions:
    * Modify the module name by replacing the default value `github.com/user/vecty-project` with your module's name in all project files.
    * Run `go mod init <modulepath>` replacing `<modulepath>` with your module's name
    * Delete the `.vscode` folder if you are not using Visual Studio Code. It is there for intellisense to work since WASM projects targets the browser, not your local computer environment. WASM projects are compiled with Go environent variables `GOOS=js` and `GOARCH=wasm`.
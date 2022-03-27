Example 2
===

Create a new article from standard input.

## Usage

1. Set access token as environment value "ESA_ACCESS_TOKEN".

    ```
    export ESA_ACCESS_TOKEN='your-access-token'
    ```

2. Run

    ```
    go run . <teamName> <postTitle>
    ```
    
3. Enter the content

    (example)

    ```
    go run . <teamName> 'sample post'
    // Please enter the content. Enter "EOF" to confirm your entry.

    # This is a sample post

    ## header 2

    - list item 1
    - list item 2
            - list item 2-1
            - list item 2-2

    EOF

    New post "sample post" is successfully created. URL is https://<teamName>.esa.io/posts/<newPostNumber>
    ```
    
    New post like following created.
    
    ![go-esa example 2](https://user-images.githubusercontent.com/9986092/160273045-568d484d-ba06-4129-879e-bc1ac6060c0c.png)
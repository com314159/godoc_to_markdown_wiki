# Parse go function doc to markdown wiki

### example

  run wiki.sh in example
  ```
  bash wiki.sh
  ```
  it will generate wiki in example and push to github wiki

### Usage

  1. first step write your function comment in markdown mode
  2. write a toml file to specify the function you need to generate into markdown wiki
  3. run main.go -toml=yourtomlfile to generate wiki

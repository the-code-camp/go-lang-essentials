## Project

You need to build a simple To-Do list app as a command line interface (cli) application. The app should:
1. provide a sub-command `add` to create a new To-Do item. The new item should be stored in a file called `todo.txt`. When saving, append `#0` to the ToDo item before adding to file. This `#` will work as a line seprator and `0` indicates that the ToDo item is not complete.
    
    ```shell
    $ ./todo add "write code for module 1"
    # after adding the todo item, cat the contents of file
    $ cat todo.txt
    # output should be "write code for module 1#0"
    ```

2. provide a sub-command `print` to print all the ToDo's in the below format, the completed items should have `✓` infront:

    ```shell
    $ ./todo print

    [ ] write code for module 1
    [✓] do something
    [ ] new item
    ```

3. provide a sub-command `mark-done` to mark an item as done:

    ```shell
    $ ./todo mark-done 3
    $ ./todo print

    [ ] write code for module 1
    [✓] do something
    [✓] new item
    ```

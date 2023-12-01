#### Earnings Tracker

This project is a simple earnings tracker program in both Go and Python. It generates an earnings summary and includes text-to-speech functionality for an interactive experience.

##### Features

- Displays daily earnings summary in the console.
- Utilizes text-to-speech functionality for an interactive experience, and generates an audio `.mp3` file.

##### Go Version

The Go version of the program uses the following external packages:

- [github.com/faiface/beep](https://github.com/faiface/beep): Audio playback and MP3 decoding library for Go.

##### Python Version

The Python version of the program uses the following external library:

- [gtts](https://pypi.org/project/gTTS/): Google Text-to-Speech library for Python.

#### How to Run

##### Go program

1. Ensure you have Go installed on your machine.
2. Run the following commands in the terminal:

    ```bash
    go run main.go
    ```

3. The program will generate a daily earnings summary and play the audio using the default system player.

##### Python program

1. Ensure you have Python installed on your machine.
2. Install the required library by running:

    ```bash
    pip install gtts
    ```

3. Run the following command in the terminal:

    ```bash
    python earn.py
    ```

4. The program will generate a daily earnings summary and play the audio using the default system player.
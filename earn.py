from gtts import gTTS
import os
import time  # Add this line to import the time module

def play_audio(file_path):
    os.system(f"start {file_path}")

def display_earnings_with_audio(name, weekly_earnings, monthly_earnings):
    message = f"{name}, you are a great technical writer and Go developer.\n"
    message += f"You earn ${weekly_earnings} weekly and ${monthly_earnings} monthly. You write and build."

    tts = gTTS(text=message, lang='en')
    audio_file_path = 'earnings.mp3'
    tts.save(audio_file_path)
    
    print("=== Earnings Summary ===")
    print(message)
    print("========================")

    play_audio(audio_file_path)

def main_with_audio():
    name = "Ekemini Samuel"
    weekly_earnings = 2000
    monthly_earnings = 10000

    while True:
        display_earnings_with_audio(name, weekly_earnings, monthly_earnings)
        time.sleep(86400)  # Sleep for 24 hours (86400 seconds) before displaying again

if __name__ == "__main__":
    main_with_audio()

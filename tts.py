import sys

import gtts
from pydub import AudioSegment
from pydub.playback import play


def main():
    while(True):
        stdin = sys.stdin.readlines()
        tts = gtts.gTTS(stdin, lang='ja')
        tts.save('sound.mp3')
        sound = AudioSegment.from_mp3('sound.mp3')
        play(sound)


if __name__ == '__main__':
    main()

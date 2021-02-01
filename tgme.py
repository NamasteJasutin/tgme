#!env python3
import sys, requests, configparser
from os.path import expanduser

cfg = configparser.ConfigParser()   
cfgfile = f"{expanduser('~')}/tgme.cfg"

def load():
    global cfg, cfgfile
    # cfg['DEFAULT'] = {
    #     'TOKEN': '1499657533:AAFTfZu3KeZ2AEU_V6zVBVoMWpLmaEWmOVY',
    #     'SENDTO': '8550799',
    # }
    # with open('tgme.cfg', 'w') as cfgfile:
    #     cfg.write(cfgfile)
    # return
    loaded = cfg.read(cfgfile)
    if len(loaded) < 1:
        print("No cfg found; I need a Telegram Token and Send To ID")
        print("Press enter to continue or CTRL+C to quit")
        input()
        print("Telegram Bot Token (ask @Botfather for one):")
        token = input()
        print("Input chat ID to send to (Howto: https://www.google.com/search?q=%22Telegram%20Chat%20ID%22&)")
        sendto = input()
        print(f"Token: {token}\nChat ID: {sendto}\n")
        save = input("Save to file? [y/n]:")
        if str(save)[0].lower() == 'y':
            cfg['DEFAULT'] = {
                'TOKEN': token,
                'SENDTO': sendto,
            }
            with open(cfgfile, 'w') as cfgfile:
                cfg.write(cfgfile)
            print("Saved to tfme.cfg")
            cfg.read(cfgfile)
        else:
            print("Cancelled. Please retry.")
            sys.exit(1)

def sendTelegram(message):
    requests.post(f"https://api.telegram.org/bot{cfg['DEFAULT']['TOKEN']}/sendMessage?chat_id={cfg['DEFAULT']['SENDTO']}&text={message}")
    return

if __name__ == "__main__":
    load()
    if sys.argv[1] == "":
        print("No message to send, input:")
        print(sys.argv[1])
    else:
        sendTelegram(sys.argv[1])
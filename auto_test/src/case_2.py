# 検証 ケース 2 速度の速い操作のテスト
# 方法 1: ブラウザ(Chrome)を開き、ログイン処理（同じID、ランダムなパスワードで送信）を行う。

import selenium
from selenium import webdriver
from selenium.webdriver.common.by import By
import time
import random
import string

# WebDriverのインスタンスを作成(Chrome)
driver = selenium.webdriver.Chrome()

try:
    driver.get("http://localhost:80")

    # フォームを送信
    id = driver.find_element(By.ID,"id")
    password = driver.find_element(By.ID,"password")

    id.send_keys("test@example.com")
    password.send_keys(''.join(random.choices(string.ascii_letters, k=5))) # ランダムな5桁の英字を生成

    driver.find_element(By.ID, "loginForm").submit()
    time.sleep(3)

except Exception as e:
    print("予期しないエラーが発生しました。", e)
finally:
    # テスト終了後、ブラウザを閉じる
    driver.quit()
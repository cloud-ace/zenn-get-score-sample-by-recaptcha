# 検証 ケース 3 高頻度のリクエストのテスト
# 方法 : ブラウザ(Chrome)を開き、ログイン処理（同じID、ランダムなパスワードで送信）を、100回繰り返す

from selenium import webdriver
from selenium.webdriver.common.by import By
import time
import random
import string

# WebDriverのインスタンスを作成(Chrome)
driver = webdriver.Chrome()

try:
    driver.get("http://localhost:80")

    # フォームを送信
    for i in range(100):

        id = driver.find_element(By.ID,"id")
        password = driver.find_element(By.ID,"password")

        id.send_keys("test@example.com")
        password.send_keys(''.join(random.choices(string.ascii_letters, k=5))) # ランダムな5桁の英字を生成

        driver.find_element(By.ID, "loginForm").submit()
        time.sleep(1)
        driver.refresh()
    
    time.sleep(50)
except Exception as e:
    print("予期しないエラーが発生しました。", e)
finally:
    # テスト終了後、ブラウザを閉じる
    driver.quit()
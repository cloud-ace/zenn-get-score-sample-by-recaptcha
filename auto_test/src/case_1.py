# 検証 ケース 1
# ログイン処理を同様のリクエスト内容で、100回繰り返す

from selenium import webdriver
from selenium.webdriver.common.by import By
import time

# WebDriverのインスタンスを作成
driver = webdriver.Chrome()

try:
    driver.get("http://localhost:80")

    id = driver.find_element(By.ID,"id")
    password = driver.find_element(By.ID,"password")

    id.send_keys("your_username")
    password.send_keys("your_password")

    # フォームを送信
    for i in range(100):
        driver.find_element(By.ID, "loginForm").submit()
    
    time.sleep(50)
except Exception as e:
    print("予期しないエラーが発生しました。", e)
finally:
    テスト終了後、ブラウザを閉じる
    driver.quit()
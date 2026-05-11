import pytest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.alert import Alert
from selenium.webdriver.support.ui import Select
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

URL_BASE = "https://the-internet.herokuapp.com"


@pytest.fixture(scope="session")
def driver():
    options = Options()

    options.add_argument("--headless=new")
    options.add_argument("--no-sandbox")
    options.add_argument("--disable-dev-shm-usage")
    options.add_argument("--remote-debugging-port=9222")
    options.add_argument("--user-data-dir=/tmp/chrome-profile")
    driver = webdriver.Chrome(options=options)
    driver.maximize_window()

    yield driver 
    driver.quit()

# LOGIN (5 TESTÓW)

def test_login_success(driver):

    wait = WebDriverWait(driver, 10)
    driver.get(URL_BASE + "/login")

    driver.find_element(By.ID, "username").send_keys("tomsmith")
    driver.find_element(By.ID, "password").send_keys("SuperSecretPassword!")

    driver.find_element(By.CSS_SELECTOR, "button").click()

    assert "secure area" in driver.page_source


def test_login_wrong_password(driver):
    driver.get(URL_BASE + "/login")
    wait = WebDriverWait(driver, 10)

    username = wait.until(EC.visibility_of_element_located((By.ID, "username")))
    password = wait.until(EC.visibility_of_element_located((By.ID, "password")))

    username.send_keys("tomsmith")
    password.send_keys("wrong")

    wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, "button"))).click()

    flash = WebDriverWait(driver, 5).until(
        EC.visibility_of_element_located((By.ID, "flash"))
    )

    assert "invalid" in flash.text.lower()


# def test_login_wrong_user(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/login")
#     driver.find_element(By.ID, "username").send_keys("wrong")
#     driver.find_element(By.ID, "password").send_keys("SuperSecretPassword!")
#     driver.find_element(By.CSS_SELECTOR, "button").click()
#     assert "invalid" in driver.page_source


# def test_login_empty(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/login")
#     driver.find_element(By.CSS_SELECTOR, "button").click()
#     assert "is required" in driver.page_source


# def test_logout(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/login")
#     driver.find_element(By.ID, "username").send_keys("tomsmith")
#     driver.find_element(By.ID, "password").send_keys("SuperSecretPassword!")
#     driver.find_element(By.CSS_SELECTOR, "button").click()
#     driver.find_element(By.LINK_TEXT, "Logout").click()
#     assert "login" in driver.current_url



# # CHECKBOXES (3 TESTY)

# def test_checkbox_1(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/checkboxes")
#     box = driver.find_elements(By.CSS_SELECTOR, "input")[0]
#     box.click()
#     assert box.is_selected()


# def test_checkbox_2(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/checkboxes")
#     box = driver.find_elements(By.CSS_SELECTOR, "input")[1]
#     assert box.is_selected()


# def test_toggle_checkbox(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/checkboxes")
#     box = driver.find_elements(By.CSS_SELECTOR, "input")[0]
#     state = box.is_selected()
#     box.click()
#     assert box.is_selected() != state


# #  DROPDOWN (3 TESTY)

# def test_dropdown_1(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/dropdown")
#     select = Select(driver.find_element(By.ID, "dropdown"))
#     select.select_by_visible_text("Option 1")
#     assert select.first_selected_option.text == "Option 1"


# def test_dropdown_2(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/dropdown")
#     select = Select(driver.find_element(By.ID, "dropdown"))
#     select.select_by_visible_text("Option 2")
#     assert select.first_selected_option.text == "Option 2"


# def test_dropdown_default(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/dropdown")
#     select = Select(driver.find_element(By.ID, "dropdown"))
#     assert "Please select" in select.first_selected_option.text


# # INPUTS (3 TESTY)

# def test_input_numbers(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/inputs")
#     input_box = driver.find_element(By.TAG_NAME, "input")
#     input_box.send_keys("123")
#     assert input_box.get_attribute("value") == "123"


# def test_clear_input(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/inputs")
#     input_box = driver.find_element(By.TAG_NAME, "input")
#     input_box.send_keys("999")
#     input_box.clear()
#     assert input_box.get_attribute("value") == ""


# def test_input_letters(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/inputs")
#     input_box = driver.find_element(By.TAG_NAME, "input")
#     input_box.send_keys("abc")
#     assert input_box.get_attribute("value") == ""


# # ALERTY (3 TESTY)

# def test_alert_accept(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/javascript_alerts")
#     driver.find_element(By.XPATH, "//button[text()='Click for JS Alert']").click()
#     Alert(driver).accept()
#     assert "successfully clicked" in driver.page_source


# def test_confirm_accept(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/javascript_alerts")
#     driver.find_element(By.XPATH, "//button[text()='Click for JS Confirm']").click()
#     Alert(driver).accept()
#     assert "Ok" in driver.page_source


# def test_confirm_dismiss(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/javascript_alerts")
#     driver.find_element(By.XPATH, "//button[text()='Click for JS Confirm']").click()
#     Alert(driver).dismiss()
#     assert "Cancel" in driver.page_source



# #  NAVIGATION (3 TESTY)


# def test_home(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE)
#     assert "The Internet" in driver.title


# def test_login_page(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/login")
#     assert "Login Page" in driver.page_source


# def test_checkboxes_page(driver):
#     wait = WebDriverWait(driver, 10)
#     driver.get(URL_BASE + "/checkboxes")
#     assert "Checkboxes" in driver.title
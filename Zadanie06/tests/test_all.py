import pytest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.alert import Alert
from selenium.webdriver.support.ui import Select
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

URL_BASE = "https://the-internet.herokuapp.com"

@pytest.fixture()
def driver():
    options = webdriver.ChromeOptions()
    options.add_argument("--disable-gpu")
    options.add_argument("--no-sandbox")
    options.add_argument("--disable-dev-shm-usage")

    driver = webdriver.Chrome(options=options)
    driver.implicitly_wait(2)
    yield driver
    driver.quit()


# LOGIN TESTY

def test_login_success(driver):
    driver.get(URL_BASE + "/login")

    driver.find_element(By.ID, "username").send_keys("tomsmith")
    driver.find_element(By.ID, "password").send_keys("SuperSecretPassword!")

    driver.find_element(By.CSS_SELECTOR, "button").click()

    assert "secure area" in driver.page_source.lower()


def test_login_wrong_password(driver):
    driver.get(URL_BASE + "/login")

    driver.find_element(By.ID, "username").send_keys("tomsmith")
    driver.find_element(By.ID, "password").send_keys("wrong")

    driver.find_element(By.CSS_SELECTOR, "button").click()

    message = WebDriverWait(driver, 5).until(
        EC.visibility_of_element_located((By.ID, "flash"))
    )

    assert "invalid" in message.text.lower()


def test_login_wrong_user(driver):
    driver.get(URL_BASE + "/login")

    driver.find_element(By.ID, "username").send_keys("wrong")
    driver.find_element(By.ID, "password").send_keys("SuperSecretPassword!")

    driver.find_element(By.ID, "login").submit()

    flash = WebDriverWait(driver, 5).until(
        EC.presence_of_element_located((By.ID, "flash"))
    )

    assert "invalid" in flash.text.lower()


def test_login_empty(driver):
    driver.get(URL_BASE + "/login")

    driver.find_element(By.CSS_SELECTOR, "button").click()

    message = WebDriverWait(driver, 5).until(
        EC.visibility_of_element_located((By.ID, "flash"))
    )

    assert "invalid" in message.text.lower()


def test_logout(driver):
    driver.get(URL_BASE + "/login")

    driver.find_element(By.ID, "username").send_keys("tomsmith")
    driver.find_element(By.ID, "password").send_keys("SuperSecretPassword!")
    driver.find_element(By.CSS_SELECTOR, "button").click()

    logout_btn = WebDriverWait(driver, 5).until(
        EC.element_to_be_clickable((By.LINK_TEXT, "Logout"))
    )

    logout_btn.click()

    WebDriverWait(driver, 5).until(
        EC.url_contains("/login")
    )

    assert "/login" in driver.current_url


# CHECKBOXES

def test_checkbox_1(driver):
    driver.get(URL_BASE + "/checkboxes")

    box = driver.find_elements(By.CSS_SELECTOR, "input")[0]

    driver.execute_script("arguments[0].click();", box)

    assert box.is_selected()


def test_checkbox_2(driver):
    driver.get(URL_BASE + "/checkboxes")

    box = driver.find_elements(By.CSS_SELECTOR, "input")[1]

    assert box.is_selected()


def test_toggle_checkbox(driver):
    driver.get(URL_BASE + "/checkboxes")

    box = driver.find_elements(By.CSS_SELECTOR, "input")[0]

    state = box.is_selected()

    driver.execute_script("arguments[0].click();", box)

    assert box.is_selected() != state



# DROPDOWN

def test_dropdown_1(driver):
    driver.get(URL_BASE + "/dropdown")

    select = Select(driver.find_element(By.ID, "dropdown"))
    select.select_by_visible_text("Option 1")

    assert select.first_selected_option.text == "Option 1"


def test_dropdown_2(driver):
    driver.get(URL_BASE + "/dropdown")

    select = Select(driver.find_element(By.ID, "dropdown"))
    select.select_by_visible_text("Option 2")

    assert select.first_selected_option.text == "Option 2"


def test_dropdown_default(driver):
    driver.get(URL_BASE + "/dropdown")

    select = Select(driver.find_element(By.ID, "dropdown"))
    assert "Please select" in select.first_selected_option.text



# INPUTS

def test_input_numbers(driver):
    driver.get(URL_BASE + "/inputs")

    input_box = driver.find_element(By.TAG_NAME, "input")
    input_box.send_keys("123")

    assert input_box.get_attribute("value") == "123"


def test_clear_input(driver):
    driver.get(URL_BASE + "/inputs")

    input_box = driver.find_element(By.TAG_NAME, "input")
    input_box.send_keys("999")
    input_box.clear()

    assert input_box.get_attribute("value") == ""


def test_input_letters(driver):
    driver.get(URL_BASE + "/inputs")

    input_box = driver.find_element(By.TAG_NAME, "input")
    input_box.send_keys("abc")

    assert input_box.get_attribute("value") == ""



# ALERTY

def test_alert_accept(driver):
    driver.get(URL_BASE + "/javascript_alerts")

    driver.find_element(By.XPATH, "//button[text()='Click for JS Alert']").click()

    alert = WebDriverWait(driver, 5).until(EC.alert_is_present())
    alert.accept()

    assert "successful" in driver.page_source.lower()


def test_confirm_accept(driver):
    driver.get(URL_BASE + "/javascript_alerts")

    driver.find_element(By.XPATH, "//button[text()='Click for JS Confirm']").click()

    alert = WebDriverWait(driver, 5).until(EC.alert_is_present())
    alert.accept()

    assert "ok" in driver.page_source.lower()


def test_confirm_dismiss(driver):
    driver.get(URL_BASE + "/javascript_alerts")

    driver.find_element(By.XPATH, "//button[text()='Click for JS Confirm']").click()

    alert = WebDriverWait(driver, 5).until(EC.alert_is_present())
    alert.dismiss()

    result = driver.find_element(By.ID, "result")

    assert "cancelled" in result.text.lower() or "cancel" in result.text.lower()


# NAVIGATION

def test_home(driver):
    driver.get(URL_BASE)
    assert "the internet" in driver.title.lower()


def test_login_page(driver):
    driver.get(URL_BASE + "/login")
    assert "login page" in driver.page_source.lower()


def test_checkboxes_page(driver):
    driver.get(URL_BASE + "/checkboxes")
    assert "checkboxes" in driver.page_source.lower()
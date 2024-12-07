from datetime import datetime, timedelta

class CommonException(Exception):
    def __init__(self, detail, status_code, request_code=None):
        self.detail = detail
        self.status_code = status_code
        self.request_code = request_code
        super().__init__(detail)

class User:
    def __init__(self, username, deleted=False, count_wrong_signin=0, password_lifetime_end=None):
        self.username = username
        self.deleted = deleted
        self.count_wrong_signin = count_wrong_signin
        self.password_lifetime_end = password_lifetime_end

def auth_conditions(user):
    now = datetime.now()

    if user.deleted:
        raise CommonException("Пользователь удален.", 403)

    if user.count_wrong_signin >= 5:
        raise CommonException("Ваша учетная запись заблокирована, обратитесь к Администратору.", 403, 207)

    if user.password_lifetime_end and now > user.password_lifetime_end:
        if now > user.password_lifetime_end + timedelta(days=90):
            raise CommonException(
                "Учётная запись заблокирована. Причина: За последние 90 дней не осуществлялся вход в систему. Обратитесь к администратору.",
                419,
                207
            )
        else:
            raise CommonException(
                "Учётная запись временно заблокирована. Причина: За последние 60 дней не осуществлялся вход в систему. Пожалуйста смените пароль.",
                423,
                207
            )

    return True

if __name__ == "__main__":
    user = User(
        username="test_user",
        deleted=False,
        count_wrong_signin=3,
        password_lifetime_end=datetime.now() - timedelta(days=91)
    )

    try:
        if auth_conditions(user):
            print("Авторизация успешна!")
    except CommonException as e:
        print(f"Ошибка авторизации: {e.detail} (Код: {e.status_code}, Запрос: {e.request_code})")

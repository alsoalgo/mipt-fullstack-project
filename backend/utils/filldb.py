import psycopg2
import os
import yaml
from datetime import date, timedelta


lotte_description = """Роскошный пятизвездочный ЛОТТЕ ОТЕЛЬ МОСКВА, расположенный на пересечении Нового Арбата и Садового кольца, открылся в сентябре 2010 года. За годы присутствия на рынке отель завоевал множество престижных международных наград и заслужил признание путешественников со всего мира."""
cosmos_description = """Легендарный «Космос» стоит на проспекте Мира — одной из главных артерий Москвы. Гостиница граничит с зеленым массивом на северо-востоке столицы, оазисе тишины и спокойствия – национальным парком «Лосиный остров»."""
delta_description = """Предлагаем к Вашим услугам 2000 комфортабельных номеров различной категории: от уютного стандарта до элегантных люксов. Размещение детей до 7 лет в номере с родителями – бесплатно. Для детей до 14 лет стоимость дополнительной кровати составляет 1000 рублей."""
mariott_description = """Пятизвездочный отель Сафмар Грандъ Москва (бывший Марриотт Гранд) открыл свои двери для гостей в 1997 году. Отель удачно вписан в архитектурный облик Тверской улицы и находится в непосредственной близости к самым красивым историческим местам Москвы и в центре культурной жизни столицы."""
movenpick_description = """Отель Mövenpick Moscow Taganskaya находится в самом сердце города Москвы, в 4 минутах ходьбы от кольцевой станции метро "Таганская". Удобное расположение позволяет посетить все самые популярные достопримечательности: Красную площадь и Кремль, Большой театр, современный парк "Зарядье", старинный район Китай-город и многие другие."""

palace_description = """Palace Bridge - уникальный отель формата «городской курорт», расположившийся в самом центре Петербурга в пешей доступности от главных достопримечательностей. Приятный интерьер отеля сочетает в себе детали старинного здания купцов Елисеевых и современные архитектурные решения в скандинавском стиле."""
astoria_description = """Непревзойденное расположение в центре города. Вид на Исаакиевский собор. Узнаваемый исторический фасад, роскошные интерьеры, камерная атмосфера. Классический послеобеденный чай в гостиной "Ротонда". История гостиницы «Астория» - это история Санкт-Петербурга."""
grand_emerald_description = """Отель расположен в аристократическом центре, который некогда назывался Рождественской частью, всего в нескольких минутах от знаменитого Невского проспекта и Московского вокзала, Смольного собора, Таврического сада и других достопримечательностей города Петра."""
four_season_lion_description = """Откройте для себя магию Санкт-Петербурга — во время индивидуальной экскурсии по легендарным достопримечательностям или вечерней прогулки на катере по каналам города. А пока вы определяетесь с выбором, позвольте нам угостить вас шампанским с икрой в духе подлинно царского гостеприимства."""
wawelberg_description = """Мы объединили в отеле особенную архитектуру, искусство, индивидуальность, тенденции и тренды в гастрономии, культуре и бизнесе, инновации и собственный театр. Всё это сложилось в богатую и многоцветную палитру, олицетворяющую современный Петербург."""


def read_db_config(env):
    with open('../configs/' + env + '/db.yml') as db_config_file:
        config = yaml.load(db_config_file, Loader=yaml.FullLoader)
    return config


def execute_commands(db_params, commands):
    conn = None
    try:
        conn = psycopg2.connect(**db_params)
        cur = conn.cursor()

        for command in commands:
            cur.execute(command)

        cur.close()
        conn.commit()
    except (Exception, psycopg2.DatabaseError) as error:
        print(error)
    finally:
        if conn is not None and not conn.closed:
            conn.close()


def clear_db(db_params):
    execute_commands(db_params, ['truncate hotel restart identity;'])
    execute_commands(db_params, ['truncate available_room restart identity;'])
    execute_commands(db_params, ['truncate popular_destination restart identity;'])
    execute_commands(db_params, ['truncate users restart identity;'])
    execute_commands(db_params, ['truncate hotel_order restart identity;'])
    execute_commands(db_params, ['truncate tokens restart identity;'])
    execute_commands(db_params, ['truncate user_question restart identity;'])


def quoted(value):
    return '\'' + value + '\''


def hotel_value(city, title, description, image_url):
    return '(' + quoted(city) + ', ' + quoted(title) + ', ' + quoted(description) + ', ' + quoted(image_url) + ')' 


def fill_hotel(db_params):
    prefix = "insert into hotel(city, title, description, image_url) values "
    hotels = [
        #city title description image_url
        ['москва', 'Лотте Отель', lotte_description, 'https://krasivodel.ru/wp-content/uploads/2020/09/Samyye-dorogiye-oteli-moskvy10.jpg'],
        ['москва', 'Космос', cosmos_description, 'https://hotelofmoscow.ru/wp-content/uploads/2017/07/49422617.jpg'],
        ['москва', 'Дельта Измайлово', delta_description, 'https://avatars.mds.yandex.net/get-altay/5484072/2a0000017ef7248331c74c19d26d319c546a/XXXL'],
        ['москва', 'Марриотт Гранд', mariott_description, 'https://photos.hotelbeds.com/giata/xxl/00/006800/006800a_hb_a_001.jpg'],
        ['москва', 'Movenpick', movenpick_description, 'https://a.travelcdn.mts.ru/property-photos/1633728227/2347900126/a4a481e310b6103b9a9964c8fc87888a779a3350.jpeg'],
        ['moscow', 'Лотте Отель', lotte_description, 'https://krasivodel.ru/wp-content/uploads/2020/09/Samyye-dorogiye-oteli-moskvy10.jpg'],
        ['moscow', 'Космос', cosmos_description, 'https://hotelofmoscow.ru/wp-content/uploads/2017/07/49422617.jpg'],
        ['moscow', 'Дельта Измайлово', delta_description, 'https://avatars.mds.yandex.net/get-altay/5484072/2a0000017ef7248331c74c19d26d319c546a/XXXL'],
        ['moscow', 'Марриотт Гранд', mariott_description, 'https://photos.hotelbeds.com/giata/xxl/00/006800/006800a_hb_a_001.jpg'],
        ['moscow', 'Movenpick', movenpick_description, 'https://a.travelcdn.mts.ru/property-photos/1633728227/2347900126/a4a481e310b6103b9a9964c8fc87888a779a3350.jpeg'],
        ['санкт-петербург', 'Palace Bridge', palace_description, 'https://thumb.tildacdn.com/tild6136-3038-4261-b634-356538376536/-/format/webp/photo.jpg'],
        ['санкт-петербург', 'Астория', astoria_description, 'https://homyrouz.ru/800/600/https/travelata-a.akamaihd.net/thumbs/1920x1080/upload/2015_03/content_hotel_56fde9811dcc55.47543375.JPEG'],
        ['санкт-петербург', 'Гранд Отель Эмеральд', grand_emerald_description, 'https://s.inyourpocket.com/gallery/238236.jpg'],
        ['санкт-петербург', 'Four Seasons Lion Palace ', four_season_lion_description, 'https://reghotel.com/wp-content/uploads/lushie-spa-oteli-rossii-chetire-sezona.jpg'],
        ['санкт-петербург', 'Wawelberg', wawelberg_description, 'https://avatars.dzeninfra.ru/get-zen_doc/8269145/pub_6422901b21067969f321e512_642d56b6b44dac4712cdc569/scale_1200'],
        ['saint-petersburg', 'Palace Bridge', palace_description, 'https://thumb.tildacdn.com/tild6136-3038-4261-b634-356538376536/-/format/webp/photo.jpg'],
        ['saint-petersburg', 'Астория', astoria_description, 'https://homyrouz.ru/800/600/https/travelata-a.akamaihd.net/thumbs/1920x1080/upload/2015_03/content_hotel_56fde9811dcc55.47543375.JPEG'],
        ['saint-petersburg', 'Гранд Отель Эмеральд', grand_emerald_description, 'https://s.inyourpocket.com/gallery/238236.jpg'],
        ['saint-petersburg', 'Four Seasons Lion Palace ', four_season_lion_description, 'https://reghotel.com/wp-content/uploads/lushie-spa-oteli-rossii-chetire-sezona.jpg'],
        ['saint-petersburg', 'Wawelberg', wawelberg_description, 'https://avatars.dzeninfra.ru/get-zen_doc/8269145/pub_6422901b21067969f321e512_642d56b6b44dac4712cdc569/scale_1200'],
    ]
    for hotel in hotels:
        execute_commands(db_params, [prefix + hotel_value(hotel[0], hotel[1], hotel[2], hotel[3]) + ';'])


def available_room_value(hotel_id, room_count, available_date):
    return '(' + str(hotel_id) + ', ' + str(room_count) + ', ' + quoted(available_date) + '::date)'


def fill_available_room(db_params):
    prefix = "insert into available_room(hotel_id, room_count, available_date) values "
    for hotel_id in range(1, 21):
        today = date.today().strftime('%Y-%m-%d')
        execute_commands(db_params, [prefix + available_room_value(hotel_id, 1, today) + ';'])
        for diff in range(1, 15):
            date_ = (date.today() + timedelta(days=diff)).strftime('%Y-%m-%d')
            execute_commands(db_params, [prefix + available_room_value(hotel_id, 1, date_) + ';'])


def popular_destination_value(city, cost, image_url):
    return '(' + quoted(city) + ', ' + str(cost) + ', ' + quoted(image_url) + '::text' + ')'


def fill_popular_destination(db_params):
    prefix = "insert into popular_destination(city, cost, image_url) values "
    destinations = [
        ['москва', 5000, 'https://sportishka.com/uploads/posts/2022-04/1650612987_17-sportishka-com-p-sovremennaya-moskva-krasivo-foto-18.jpg'],
        ['казань', 3000, 'https://ic.pics.livejournal.com/zdorovs/16627846/1551759/1551759_original.jpg'],
        ['новосибирск', 4000, 'https://sdelanounas.ru/i/a/w/1/f_aW1nLmdlbGlvcGhvdG8uY29tL25zazIwMjB3LzI5X25zazIwMjB3LmpwZz9fX2lkPTEzMTAxMQ==.jpeg'],
        ['петрозаводск', 2000, 'https://mykaleidoscope.ru/x/uploads/posts/2022-09/1663251393_9-mykaleidoscope-ru-p-stolitsa-karelii-petrozavodsk-instagram-9.jpg'],
        ['тверь',3000, 'https://levelvan.ru/upload/tmp/62cc755759c71.jpg'],
        ['хабаровск', 6000, 'https://laguna-27.ru/wp-content/uploads/2021/04/upload-717804a0-a2a3-11e8-a984-65405bff09b2-1500x1000.jpg'],
    ]
    for destination in destinations:
        execute_commands(db_params, [prefix + popular_destination_value(destination[0], destination[1], destination[2]) + ';'])


def fill_db(db_params):
    fill_hotel(db_params)
    fill_available_room(db_params)
    fill_popular_destination(db_params)


def main():
    env = os.environ.get('ENV', 'dev')
    config = read_db_config(env)
    db_params = config['postgres']
    clear_db(db_params)
    fill_db(db_params)


if __name__ == '__main__':
    main()
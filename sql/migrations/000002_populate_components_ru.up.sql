BEGIN;

INSERT INTO
    languages (name, label)
VALUES
    ('Русский', 'RU');

INSERT INTO
    panel_measurement_units (name, language_id)
VALUES
    (
        'мкмоль/л',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'ед/л',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'расчетный показатель',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'ммоль/л',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'г/л',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'ед.pH',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'расчетный, г/л',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'расчетный, ммоль/л',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'расчетный, мг/дл',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'x10^12/л',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'x10^9/л',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'мм/ч',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'в п/зр (HPF)',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'мкм^3 (фл)',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'пг',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'на 100 лейкоцитов',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        '%',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    );

INSERT INTO
    panel_components (
        name,
        unit_id,
        min_amount,
        max_amount,
        comment,
        language_id
    )
VALUES
    (
        'Билирубин общий (TBil)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'мкмоль/л'
        ),
        null,
        10.0,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Билирубин прямой (DBil)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'мкмоль/л'
        ),
        null,
        5.5,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'АСТ (GOT)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ед/л'
        ),
        12,
        45,
        'возраст старше 6 месяцев',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'АЛТ (GPT)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ед/л'
        ),
        18,
        60,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Коэффициент Ритиса',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'расчетный показатель'
        ),
        1.1,
        1.3,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Мочевина (Urea)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ммоль/л'
        ),
        5.4,
        12.1,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Креатинин (Creat)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'мкмоль/л'
        ),
        70,
        165,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Общий белок (Prot, total)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'г/л'
        ),
        54,
        78,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Альбумин (Alb)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'г/л'
        ),
        24,
        38,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Щелочная фосфотаза (ALP, IFCC)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'г/л'
        ),
        null,
        55,
        'возраст старше 6 месяцев',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Альфа-Амилаза, общая (ά-Amylase,total)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'г/л'
        ),
        500,
        1200,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Глюкоза (Glu)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ммоль/л'
        ),
        3.3,
        6.8,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'ЛДГ (LDH, IFCC)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ед/л'
        ),
        35,
        500,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'ГГТ (γ-GT)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ед/л'
        ),
        0,
        4,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Холестерин (Chol, total)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ммоль/л'
        ),
        1.90,
        3.90,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Триглицериды (Trig)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ммоль/л'
        ),
        0.38,
        1.10,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'КФК (CK)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ед/л'
        ),
        150,
        350,
        'возраст старше 12 месяцев',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Калий (Potassium)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ммоль/л'
        ),
        3.6,
        5.5,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Натрий (Sodium)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ммоль/л'
        ),
        144,
        158,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Фосфор (Phosphate, inorg)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ммоль/л'
        ),
        1.10,
        2.30,
        'возраст старше 6 месяцев',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Кальций общий (Ca, total)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ммоль/л'
        ),
        2.00,
        2.70,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Железо (Fe)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'мкмоль/л'
        ),
        12.0,
        39.0,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Магний (Mg)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ммоль/л'
        ),
        0.8,
        1.2,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Хлор (Chloride)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ммоль/л'
        ),
        107,
        129,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Кислотность (pH)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'ед.pH'
        ),
        7.35,
        7.50,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Альбумин/глобулин (A/G ratio)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'расчетный, г/л'
        ),
        0.7,
        1.4,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Глобулин (Glob)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'г/л'
        ),
        29,
        55,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Соотношение Ca/P',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'расчетный, ммоль/л'
        ),
        1.6,
        2.3,
        'возраст старше 10 месяцев',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Соотношение Na/K',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'расчетный, ммоль/л'
        ),
        26,
        55,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Соотношение Мочевина/Креатинин (B/C ratio)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'расчетный, мг/дл'
        ),
        null,
        null,
        'при азотемии: > 30 пре-/постренальная; <20 ренальная А.  или другая причина',
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Гематокрит (Hct, PCV)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        29,
        48,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Гемоглобин (Hgb)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'г/л'
        ),
        90,
        150,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Эритроциты (RBC)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'x10^12/л'
        ),
        5.6,
        10.0,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Лейкоциты (WBC)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'x10^9/л'
        ),
        5.5,
        18.5,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Бластные клетки',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        0,
        0,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Миелоциты',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        0,
        0,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Метамиелоциты',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        0,
        0,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Палочкоядные нейтрофилы (Bands)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        0,
        3,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Сегментоядерные нейтрофилы (Segs)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        35,
        75,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Эозинофилы (Eos)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        0,
        6,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Моноциты',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        1,
        4,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Базофилы (Bas)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        0,
        1,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Лимфоциты (Lym)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        25,
        55,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Тромбоциты (Plt)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'x10^9/л'
        ),
        160,
        630,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Количество тромбоцитов в п/зр',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'в п/зр (HPF)'
        ),
        8,
        30,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'СОЭ',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'мм/ч'
        ),
        1,
        13,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Количество тромбоцитов в п/зр',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'в п/зр (HPF)'
        ),
        8,
        30,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Ядерные эритроциты (нормобласты, nRBC)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'на 100 лейкоцитов'
        ),
        0,
        0,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Анизоцитоз эритроцитов (RDW)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        null,
        22,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Средняя конц. Hb в эритроците (MCHC)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = '%'
        ),
        29,
        36,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Средний объем эритроцита (MCV)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'мкм^3 (фл)'
        ),
        39,
        53,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Сред содержание Hb в эритроците (MCH)',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'пг'
        ),
        12.5,
        17.5,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Скорректир (истинные) лейкоциты',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'x10^9/л'
        ),
        6.5,
        6.5,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Палочкоядерные нейтрофилы ABS',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'x10^9/л'
        ),
        0,
        0.3,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Сегментоядерные нейтрофилы ABS',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'x10^9/л'
        ),
        2.5,
        12.50,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Эозинофилы ABS',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'x10^9/л'
        ),
        0.1,
        1.5,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Базофилы ABS',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'x10^9/л'
        ),
        0,
        0.1,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Моноциты ABS',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'x10^9/л'
        ),
        0,
        0.9,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Лимфоциты ABS',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'x10^9/л'
        ),
        1.5,
        7,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    ),
    (
        'Бластные клетки ABS',
        (
            SELECT
                id
            FROM
                panel_measurement_units
            WHERE
                name = 'x10^9/л'
        ),
        0,
        0,
        null,
        (
            SELECT
                id
            FROM
                languages
            WHERE
                name = 'Русский'
        )
    );

COMMIT;

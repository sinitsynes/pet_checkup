BEGIN;

INSERT INTO
    measurement_unit (name)
VALUES
    ('мкмоль/л'),
    ('ед/л'),
    ('расчетный показатель'),
    ('ммоль/л'),
    ('г/л'),
    ('ед.pH'),
    ('расчетный, г/л'),
    ('расчетный, ммоль/л'),
    ('расчетный, мг/дл'),
    ('x10^12/л'),
    ('x10^9/л'),
    ('мм/ч'),
    ('в п/зр (HPF)'),
    ('мкм^3 (фл)'),
    ('пг'),
    ('на 100 лейкоцитов'),
    ('%');

INSERT INTO
    component (
        name,
        unit_id,
        min_amount,
        max_amount,
        comment
    )
VALUES
    (
        'Билирубин общий (TBil)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'мкмоль/л'
        ),
        null,
        10.0,
        null
    ),
    (
        'Билирубин прямой (DBil)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'мкмоль/л'
        ),
        null,
        5.5,
        null
    ),
    (
        'АСТ (GOT)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ед/л'
        ),
        12,
        45,
        'возраст старше 6 месяцев'
    ),
    (
        'АЛТ (GPT)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ед/л'
        ),
        18,
        60,
        null
    ),
    (
        'Коэффициент Ритиса',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'расчетный показатель'
        ),
        1.1,
        1.3,
        null),
    (
        'Мочевина (Urea)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ммоль/л'
        ),
        5.4,
        12.1,
        null),
    (
        'Креатинин (Creat)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'мкмоль/л'
        ),
        70,
        165,
        null
    ),
    (
        'Общий белок (Prot, total)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'г/л'
        ),
        54,
        78,
        null),
    (
        'Альбумин (Alb)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'г/л'
        ),
        24,
        38,
        null
    ),
    (
        'Щелочная фосфотаза (ALP, IFCC)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'г/л'
        ),
        null,
        55,
        'возраст старше 6 месяцев'
    ),
    (
        'Альфа-Амилаза, общая (ά-Amylase,total)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'г/л'
        ),
        500,
        1200,
        null
    ),
    (
        'Глюкоза (Glu)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ммоль/л'
        ),
        3.3,
        6.8,
        null
    ),
    (
        'ЛДГ (LDH, IFCC)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ед/л'
        ),
        35,
        500,
        null
    ),
    (
        'ГГТ (γ-GT)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ед/л'
        ),
        0,
        4,
        null
    ),
    (
        'Холестерин (Chol, total)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ммоль/л'
        ),
        1.90,
        3.90,
        null
    ),
    (
        'Триглицериды (Trig)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ммоль/л'
        ),
        0.38,
        1.10,
        null
    ),
    (
        'КФК (CK)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ед/л'
        ),
        150,
        350,
        'возраст старше 12 месяцев'
    ),
    (
        'Калий (Potassium)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ммоль/л'
        ),
        3.6,
        5.5,
        null
    ),
    (
        'Натрий (Sodium)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ммоль/л'
        ),
        144,
        158,
        null
    ),
    (
        'Фосфор (Phosphate, inorg)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ммоль/л'
        ),
        1.10,
        2.30,
        'возраст старше 6 месяцев'
    ),
    (
        'Кальций общий (Ca, total)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ммоль/л'
        ),
        2.00,
        2.70,
        null
    ),
    (
        'Железо (Fe)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'мкмоль/л'
        ),
        12.0,
        39.0,
        null
    ),
    (
        'Магний (Mg)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ммоль/л'
        ),
        0.8,
        1.2,
        null
    ),
    (
        'Хлор (Chloride)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ммоль/л'
        ),
        107,
        129,
        null
    ),
    (
        'Кислотность (pH)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'ед.pH'
        ),
        7.35,
        7.50,
        null
    ),
    (
        'Альбумин/глобулин (A/G ratio)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'расчетный, г/л'
        ),
        0.7,
        1.4,
        null
    ),
    (
        'Глобулин (Glob)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'г/л'
        ),
        29,
        55,
        null
    ),
    (
        'Соотношение Ca/P',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'расчетный, ммоль/л'
        ),
        1.6,
        2.3,
        'возраст старше 10 месяцев'
    ),
    (
        'Соотношение Na/K',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'расчетный, ммоль/л'
        ),
        26,
        55,
        null
    ),
    (
        'Соотношение Мочевина/Креатинин (B/C ratio)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'расчетный, мг/дл'
        ),
        null,
        null,
        'при азотемии: > 30 пре-/постренальная; <20 ренальная А.  или другая причина'
    ),
    (
        'Гематокрит (Hct, PCV)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        29,
        48,
        null
    ),
    (
        'Гемоглобин (Hgb)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'г/л'
        ),
        90,
        150,
        null
    ),
    (
        'Эритроциты (RBC)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'x10^12/л'
        ),
        5.6,
        10.0,
        null
    ),
    (
        'Лейкоциты (WBC)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'x10^9/л'
        ),
        5.5,
        18.5,
        null
    ),
    (
        'Бластные клетки',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        0,
        0,
        null
    ),
    (
        'Миелоциты',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        0,
        0,
        null
    ),
    (
        'Метамиелоциты',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        0,
        0,
        null
    ),
    (
        'Палочкоядные нейтрофилы (Bands)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        0,
        3,
        null
    ),
    (
        'Сегментоядерные нейтрофилы (Segs)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        35,
        75,
        null
    ),
    (
        'Эозинофилы (Eos)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        0,
        6,
        null
    ),
    (
        'Моноциты',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        1,
        4,
        null
    ),
    (
        'Базофилы (Bas)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        0,
        1,
        null
    ),
    (
        'Лимфоциты (Lym)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        25,
        55,
        null
    ),
    (
        'Тромбоциты (Plt)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'x10^9/л'
        ),
        160,
        630,
        null
    ),
    (
        'Количество тромбоцитов в п/зр',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'в п/зр (HPF)'
        ),
        8,
        30,
        null
    ),
    (
        'СОЭ',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'мм/ч'
        ),
        1,
        13,
        null
    ),
    (
        'Количество тромбоцитов в п/зр',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'в п/зр (HPF)'
        ),
        8,
        30,
        null
    ),
    (
        'Ядерные эритроциты (нормобласты, nRBC)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'на 100 лейкоцитов'
        ),
        0,
        0,
        null
    ),
    (
        'Анизоцитоз эритроцитов (RDW)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        null,
        22,
        null
    ),
    (
        'Средняя конц. Hb в эритроците (MCHC)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = '%'
        ),
        29,
        36,
        null
    ),
    (
        'Средний объем эритроцита (MCV)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'мкм^3 (фл)'
        ),
        39,
        53,
        null
    ),
    (
        'Сред содержание Hb в эритроците (MCH)',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'пг'
        ),
        12.5,
        17.5,
        null
    ),
    (
        'Скорректир (истинные) лейкоциты',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'x10^9/л'
        ),
        6.5,
        6.5,
        null
    ),
    (
        'Палочкоядерные нейтрофилы ABS',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'x10^9/л'
        ),
        0,
        0.3,
        null
    ),
    (
        'Сегментоядерные нейтрофилы ABS',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'x10^9/л'
        ),
        2.5,
        12.50,
        null
    ),
    (
        'Эозинофилы ABS',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'x10^9/л'
        ),
        0.1,
        1.5,
        null
    ),
    (
        'Базофилы ABS',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'x10^9/л'
        ),
        0,
        0.1,
        null
    ),
    (
        'Моноциты ABS',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'x10^9/л'
        ),
        0,
        0.9,
        null
    ),
    (
        'Лимфоциты ABS',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'x10^9/л'
        ),
        1.5,
        7,
        null
    ),
    (
        'Бластные клетки ABS',
        (
            SELECT
                id
            FROM
                measurement_unit
            WHERE
                name = 'x10^9/л'
        ),
        0,
        0,
        null
    );

COMMIT;

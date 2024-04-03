# Nöbet Planlama Hizmeti

Kullanıcıların vardiya dönemlerini yönetmelerini ve vardiya listelerini oluşturmalarını sağlayan hizmettir. Bu hizmet, kullanıcıların belirli bir tarih aralığında hangi gün ve saatlerde görevde olacaklarını seçmelerine ve vardiyalarını oluşturmalarına olanak tanır. Buna ek olarak, kullanıcılar vardiya dönemlerine ait vardiya listesini ve geçici değişiklik taleplerini görüntüleyebilir ve mevcut vardiyaları silme veya güncelleme seçeneğine sahip olabilirler. Bildirimler aracılığıyla kullanıcılar vardiya dönemleriyle ilgili güncellemelerden haberdar edilebilmektedir. Vardiya Hizmeti, kullanıcıların vardiyalarını düzenli bir şekilde planlamalarına ve vardiya listelerini yönetmelerine yardımcı olan kullanıcı dostu bir platform sağlar.

## Özellikler

- CRUD işlemleri için RESTful API uç noktaları.
- JWT Kimlik Doğrulama. (**Daha sonra eklenecek**)
- Rate Limiting.
- Swagger Dokümantasyonu.
- GORM kullanarak PostgreSQL veritabanı entegrasyonu.
- Redis önbelleği.
- Kolay kurulum ve dağıtım için Dockerized uygulama.
- Grafana, Prometheus ve Jaeger

## İçerik

- [Hızlı başlangıç](#hızlı-başlangıç)
- [Proje yapısı](#proje-yapısı)

## Hızlı başlangıç

Bu Nöbet Planlama servisini Docker ile çalıştırabiliriz. Burada, bu projeyi çalıştırmak için her iki yolu da sunuyorum.

- Bu projeyi klonlayın

```bash
# Çalışma alanınıza gidin
cd your-workspace

# Bu projeyi çalışma alanınıza klonlayın
git clone https://github.com/YunusEmreAlps/shift-scheduler-service.git

# Proje kök dizinine taşı
cd shift-scheduler-service
```

### Docker ile çalıştır

- Yapılandırmanızla birlikte kök dizinde `.env.example` benzeri bir `.env` dosyası oluşturun.
- Docker ve Docker Compose yükleyin.
- Make` veya `docker` komutlarını çalıştırın.

```bash
make local // tüm kapsayıcıları çalıştır
make run // hata ayıklayıcı eklemek veya projeyi yeniden oluşturmak/çalıştırmak daha kolay bir yoldur
```

```bash
docker-compose -f docker-compose.local.yml up
```

- API'ye `http://localhost:9097` kullanarak erişin

#### Docker geliştirme kullanımı

```bash
make docker
```

### SWAGGER UI

```bash
https://localhost:5000/swagger/index.html
```

### Jaeger Kullanıcı Arayüzü

```bash
http://localhost:16686
```

### Prometheus Kullanıcı Arayüzü

```bash
http://localhost:9090
```

### Grafana Kullanıcı Arayüzü

```bash
http://localhost:3000
```

Grafana'yı farklı bir şekilde yapılandırmadıysanız, varsayılan olarak [localhost](http://localhost:3000) kullanacak şekilde ayarlanmıştır. Oturum açma sayfasında, kullanıcı adı ve parola için **admin** girin.

## Proje yapısı

```bash
shift-scheduler-service/
|-- assets/
|   |-- v1_db.png
|-- config/
|   |-- config.go
|   |-- sample.env.yaml
|-- docker/
|-- internal/
|   |-- handlers/
|       |-- handlers.go
|   |-- middleware/
|       |-- jwt_middleware.go
|   |-- models/
|       |-- shift_schedule.go
|       |-- jsonb.go
|       |-- pagination.go
|       |-- jwt_response.go
|       |-- s3.go
|-- pkg/
|   |-- db/
|       |-- aws/
|           |-- aws.go
|       |-- postgres/
|           |-- db_conn.go
|       |-- redis/
|           |-- conn.go
|   |-- httpErrors/
|       |-- http_errors.go
|   |-- logger/
|       |-- logger.go
|   |-- metric/
|       |-- metric.go
|   |-- postman/
|       |-- Shift_Scheduler_Service.postman_collection.json
|   |-- sanitize/
|       |-- sanitize.go
|   |-- utils/
|       |-- helpers.go
|       |-- http.go
|       |-- pagination.go
|       |-- validator.go
|-- Dockerfile
|-- go.mod
|-- go.sum
|-- main.go
|-- Makefile
|-- README.md
```

### `assets`

Kodunuzla birlikte kullanılacak diğer varlıklar (resimler, logolar, vb.).

### `config`

Yapılandırma. Önce `config.yml` okunur, ardından ortam değişkenleri eşleşirse yaml yapılandırmasının üzerine yazılır.
Yapılandırma yapısı `config.go` içerisindedir.
Env-required: true` etiketi bir değer belirtmenizi zorunlu kılar (yaml ya da ortam değişkenlerinde).

Yapılandırmayı yaml'dan okumak 12 faktör ideolojisiyle çelişir, ancak pratikte
tüm yapılandırmayı ENV'den okumak.
Varsayılan değerlerin yaml'da olduğu ve güvenliğe duyarlı değişkenlerin ENV'de tanımlandığı varsayılır.

### `docs`

Swagger belgeleri. [swag](https://github.com/swaggo/swag) kütüphanesi tarafından otomatik olarak oluşturulur.
Hiçbir şeyi kendiniz düzeltmenize gerek yok.

### `/internal`

Özel uygulama ve kütüphane kodu. Bu, başkalarının kendi uygulamalarında veya kitaplıklarında içe aktarmasını istemediğiniz koddur. Bu düzen modelinin Go derleyicisinin kendisi tarafından zorlandığını unutmayın. Daha fazla ayrıntı için Go 1.4 sürüm notlarına bakın. En üst düzey dahili dizinle sınırlı olmadığınızı unutmayın. Proje ağacınızın herhangi bir seviyesinde birden fazla dahili dizine sahip olabilirsiniz.

İsteğe bağlı olarak, paylaşılan ve paylaşılmayan dahili kodunuzu ayırmak için dahili paketlerinize biraz ekstra yapı ekleyebilirsiniz. Bu gerekli değildir (özellikle küçük projeler için), ancak amaçlanan paket kullanımını gösteren görsel ipuçlarına sahip olmak güzeldir. Gerçek uygulama kodunuz /internal/app dizinine (örneğin, /internal/app/myapp) ve bu uygulamalar tarafından paylaşılan kod /internal/pkg dizinine (örneğin, /internal/pkg/myprivlib) gidebilir.

### `/pkg`

Harici uygulamalar tarafından kullanılması uygun olan kütüphane kodu (örneğin, /pkg/mypubliclib). Diğer projeler bu kütüphanelerin çalışmasını bekleyerek içe aktaracaktır, bu nedenle buraya bir şey koymadan önce iki kez düşünün :-) Dahili dizinin özel paketlerinizin içe aktarılmamasını sağlamak için daha iyi bir yol olduğunu unutmayın, çünkü Go tarafından zorlanmaktadır. pkg dizini, bu dizindeki kodun başkaları tarafından kullanılmasının güvenli olduğunu açıkça bildirmek için hala iyi bir yoldur. Travis Jeffery tarafından yazılan I'll take pkg over internal blog yazısı, pkg ve internal dizinleri ve bunları kullanmanın ne zaman mantıklı olabileceği hakkında iyi bir genel bakış sağlar.

Ayrıca, kök dizininiz Go olmayan birçok bileşen ve dizin içerdiğinde Go kodunu tek bir yerde gruplamanın bir yoludur ve çeşitli Go araçlarını çalıştırmayı kolaylaştırır (bu konuşmalarda belirtildiği gibi: GopherCon EU 2018'den Endüstriyel Programlama için En İyi Uygulamalar, GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps ve GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go).

Hangi popüler Go depolarının bu proje düzen modelini kullandığını görmek istiyorsanız /pkg dizinine bakın. Bu yaygın bir düzen modelidir, ancak evrensel olarak kabul görmez ve Go topluluğundaki bazıları bunu önermez.

Uygulama projeniz gerçekten küçükse ve fazladan bir iç içe geçme seviyesi çok fazla değer katmıyorsa (gerçekten istemiyorsanız :-)) kullanmamanızda sorun yoktur. Yeterince büyüdüğünde ve kök dizininiz oldukça meşgul olduğunda düşünün (özellikle çok sayıda Go dışı uygulama bileşeniniz varsa).

pkg dizini kökenleri: Eski Go kaynak kodu, paketleri için pkg kullanırdı ve daha sonra topluluktaki çeşitli Go projeleri bu modeli kopyalamaya başladı (daha fazla bağlam için Brad Fitzpatrick'in tweet'ine bakın).

#### `handlers`

İşleyiciler dizini, proje için ana işleyicileri veya denetleyicileri içerir. Bu işleyiciler gelen istekleri ele alır, gerekli eylemleri gerçekleştirir ve uygun yanıtları döndürür. İş mantığını kapsüllerler ve hizmetler ve veri havuzları gibi projenin diğer bileşenleriyle etkileşime girerler.

Burada açıklanan proje yapısının gerçek projede bulunan tüm dizinleri ve dosyaları içermeyebileceğini unutmamak önemlidir. Sağlanan genel bakış, projenin yapısını ve organizasyonunu anlamakla ilgili temel dizinlere odaklanmaktadır.

## Bu projede kullanılan başlıca paketler

- **gin**: Gin, Go (Golang) ile yazılmış bir HTTP web çerçevesidir. Çok daha iyi performansa sahip Martini benzeri bir API'ye sahiptir - 40 kata kadar daha hızlı. Müthiş bir performansa ihtiyacınız varsa, kendinize biraz Gin alın.
- **gorm**: Golang için Nesne İlişkisel Eşleme (ORM) kütüphanesi. ORM, nesne yönelimli programlama dillerini kullanarak uyumsuz tip sistemleri arasında veri dönüştürür.
- **postgreSQL go sürücüsü**: PostgreSQL için Resmi Golang sürücüsü.
- **jwt**: JSON Web Belirteçleri, talepleri iki taraf arasında güvenli bir şekilde temsil etmek için açık, endüstri standardı bir RFC 7519 yöntemidir. Erişim Belirteci ve Yenileme Belirteci için kullanılır.
- **viper**: .env` dosyasından yapılandırma yüklemek için. Fangs ile yapılandırmaya gidin. JSON, TOML, YAML, HCL, INI, envfile veya Java özellikleri biçimlerinde bir yapılandırma dosyasını bulun, yükleyin ve açın.
- bcrypt**: bcrypt paketi Provos ve Mazières'in bcrypt uyarlanabilir karma algoritmasını uygular.
- testify**: Standart kütüphane ile güzel bir şekilde oynayan ortak assertions ve mocks içeren bir araç seti.
- Daha fazla paketi `go.mod` içinde kontrol edin.

## Veritabanı Diyagramı

![DB](/assets/v1_db.png)


## JWT Kimlik Doğrulama Ara Katmanı olmadan Genel API İstek Akışı

![Public API Request Flow](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-arch-public-api-request-flow.png?raw=true)

## JWT Kimlik Doğrulama Ara Katmanı ile Özel API İstek Akışı

![Private API Request Flow](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-arch-private-api-request-flow.png?raw=true)

### Örnek API İsteği ve Yanıtı

pkg/postman içerisinde yer alan dosyayı postman uygulamasına sürükleyin. 

- **health-check**

  - İstek

  ```bash
    curl --location 'http://localhost:9097/shift-scheduler-service/health'
  ```

  - Yanıt

  ```json
  {
      "message": "Shift Scheduler Service 1.0.0 is running on port 9097.",
      "status": true
  }
  ```

- **shift-schedules**

  - İstek

  ```bash
    curl --location 'http://localhost:9097/shift-scheduler-service/shift-schedules'
  ```

  - Yanıt

  ```json
  {
    "status": true,
    "intent": "cld:::shiftservice:::/shift-schedules",
    "message": [
        {
            "ID": 0,
            "alias": "Shift 1",
            "description": "Description 1",
            "frequency": 7,
            "start_date": "2023-01-01T00:00:00+03:00",
            "end_date": "2023-10-01T00:00:00+03:00",
            "year": 2023,
            "status": 0,
            "organization": [
                {
                    "id": 0,
                    "name": "Group 1",
                    "phone": "",
                    "mail": "",
                    "description": ""
                }
            ],
            "manager": [
                {
                    "id": 0,
                    "name": "Manager 1",
                    "phone": "",
                    "mail": "",
                    "description": ""
                }
            ],
            "shifts": [
                {
                    "id": 0,
                    "start": "2023-01-01 00:00:00",
                    "end": "2023-02-01 00:00:00",
                    "user": {
                        "name": "User 1",
                        "mail": "",
                        "phone": ""
                    }
                },
                {
                    "id": 1,
                    "start": "2023-01-01 00:00:00",
                    "end": "2023-02-01 00:00:00",
                    "user": {
                        "name": "User 2",
                        "mail": "",
                        "phone": ""
                    }
                },
                {
                    "id": 2,
                    "start": "2023-01-01 00:00:00",
                    "end": "2023-02-01 00:00:00",
                    "user": {
                        "name": "User 3",
                        "mail": "",
                        "phone": ""
                    }
                }
            ],
            "CreatedAt": "2024-01-03T15:05:09.503045+03:00",
            "UpdatedAt": "2024-01-03T15:05:09.503045+03:00",
            "DeletedAt": null,
        }
    ]
  }
  ```

- **shift-schedules/id**

  - İstek

  ```bash
    curl --location 'http://localhost:9097/shift-scheduler-service/shift-schedules/1'
  ```

  - Yanıt

  ```json
  {
    "status": true,
    "intent": "cld:::shiftservice:::/shift-schedules/:id",
    "message": {
        "ID": 1,
        "alias": "Shift 1",
        "description": "Description 1",
        "frequency": 7,
        "start_date": "2023-01-01T00:00:00+03:00",
        "end_date": "2023-10-01T00:00:00+03:00",
        "year": 2023,
        "status": 0,
        "organization": [
            {
                "id": 0,
                "name": "Group 1",
                "phone": "",
                "mail": "",
                "description": ""
            }
        ],
        "manager": [
            {
                "id": 0,
                "name": "Manager 1",
                "phone": "",
                "mail": "",
                "description": ""
            }
        ],
        "shifts": [
            {
                "id": 0,
                "start": "2023-01-01 00:00:00",
                "end": "2023-02-01 00:00:00",
                "user": {
                    "name": "User 1",
                    "mail": "",
                    "phone": ""
                }
            },
            {
                "id": 1,
                "start": "2023-01-01 00:00:00",
                "end": "2023-02-01 00:00:00",
                "user": {
                    "name": "User 2",
                    "mail": "",
                    "phone": ""
                }
            },
            {
                "id": 2,
                "start": "2023-01-01 00:00:00",
                "end": "2023-02-01 00:00:00",
                "user": {
                    "name": "User 3",
                    "mail": "",
                    "phone": ""
                }
            }
        ],
        "CreatedAt": "2024-01-03T15:05:09.503045+03:00",
        "UpdatedAt": "2024-01-03T15:05:09.503045+03:00",
        "DeletedAt": null
    }
  }
  ```

- **shift-schedules/deleted**

  - İstek

  ```bash
    curl --location 'http://localhost:9097/shift-scheduler-service/shift-schedules/deleted'
  ```

  - Yanıt

  ```json
  {
    "status": true,
    "intent": "cld:::shiftservice:::/shift-schedules/:deleted",
    "message": {
        "ID": 2,
        "alias": "Shift 2",
        "description": "Description 1",
        "frequency": 7,
        "start_date": "2023-01-01T03:00:00+03:00",
        "end_date": "2023-11-01T03:00:00+03:00",
        "year": 2023,
        "status": 0,
        "organization": [
            {
                "id": 0,
                "name": "Group 1",
                "phone": "",
                "mail": "",
                "description": ""
            }
        ],
        "manager": [
            {
                "id": 0,
                "name": "Manager 1",
                "phone": "",
                "mail": "",
                "description": ""
            }
        ],
        "shifts": [
          {
              "id": 0,
              "start": "2023-01-01 00:00:00",
              "end": "2023-02-01 00:00:00",
              "user": {
                  "name": "User 1",
                  "mail": "",
                  "phone": ""
              },
          },
          {
              "id": 1,
              "start": "2023-01-01 00:00:00",
              "end": "2023-02-01 00:00:00",
              "user": {
                  "name": "User 2",
                  "mail": "",
                  "phone": ""
              },
          },
          {
              "id": 2,
              "start": "2023-01-01 00:00:00",
              "end": "2023-02-01 00:00:00",
              "user": {
                  "name": "User 3",
                  "mail": "",
                  "phone": ""
              },
          }
        ],
        "CreatedAt": "2023-12-28T14:31:06.91924+03:00",
        "UpdatedAt": "2023-12-28T14:42:47.283204+03:00",
        "DeletedAt": null
    }
  }  
  ```

- **shift-schedules/paginated**

  - İstek

  ```bash
    curl --location 'http://localhost:9097/shift-scheduler-service/shift-schedules/paginated?page=1&size=10&orderBy=desc'
  ```

  - Yanıt

  ```json
  {
    "status": true,
    "intent": "cld:::shiftservice:::/shift-schedules/pagination",
    "message": [
        {
            "ID": 1,
            "alias": "Shift 1",
            "description": "Description 1",
            "frequency": 7,
            "start_date": "2023-01-01T00:00:00+03:00",
            "end_date": "2023-10-01T00:00:00+03:00",
            "year": 2023,
            "status": 0,
            "organization": [
                {
                  "id": 0,
                  "name": "Group 1",
                  "phone": "",
                  "mail": "",
                  "description": ""
                }
            ],
            "manager": [
                {
                  "id": 0,
                  "name": "Manager 1",
                  "phone": "",
                  "mail": "",
                  "description": ""
                }
            ],
            "shifts": [
                {
                    "id": 0,
                    "start": "2023-01-01 00:00:00",
                    "end": "2023-02-01 00:00:00",
                    "user": {
                        "name": "User 1",
                        "mail": "",
                        "phone": ""
                    }
                },
                {
                    "id": 1,
                    "start": "2023-01-01 00:00:00",
                    "end": "2023-02-01 00:00:00",
                    "user": {
                        "name": "User 2",
                        "mail": "",
                        "phone": ""
                    }
                },
                {
                    "id": 2,
                    "start": "2023-01-01 00:00:00",
                    "end": "2023-02-01 00:00:00",
                    "user": {
                        "name": "User 3",
                        "mail": "",
                        "phone": ""
                    }
                }
            ],
            "CreatedAt": "2024-01-03T15:05:09.503045+03:00",
            "UpdatedAt": "2024-01-03T15:05:09.503045+03:00",
            "DeletedAt": null
        }
    ]
  }
  ```

- **shift-schedules/year/2024**

  - İstek

  ```bash
    curl --location 'http://localhost:9097/shift-scheduler-service/shift-schedules/year/2023'
  ```

  - Yanıt

  ```json
   {
    "status": true,
    "intent": "cld:::shiftservice:::/shift-schedules/year/:year",
    "message": {
        "ID": 1,
        "alias": "Shift 1",
        "description": "Description 1",
        "frequency": 7,
        "start_date": "2023-01-01T00:00:00+03:00",
        "end_date": "2023-10-01T00:00:00+03:00",
        "year": 2023,
        "status": 0,
        "organization": [
            {
                "id": 0,
                "name": "Group 1",
                "phone": "",
                "mail": "",
                "description": ""
            }
        ],
        "manager": [
            {
                "id": 0,
                "name": "Manager 1",
                "phone": "",
                "mail": "",
                "description": ""
            }
        ],
        "shifts": [
            {
                "id": 0,
                "start": "2023-01-01 00:00:00",
                "end": "2023-02-01 00:00:00",
                "user": {
                    "name": "User 1",
                    "mail": "",
                    "phone": ""
                }
            },
            {
                "id": 1,
                "start": "2023-01-01 00:00:00",
                "end": "2023-02-01 00:00:00",
                "user": {
                    "name": "User 2",
                    "mail": "",
                    "phone": ""
                }
            },
            {
                "id": 2,
                "start": "2023-01-01 00:00:00",
                "end": "2023-02-01 00:00:00",
                "user": {
                    "name": "User 3",
                    "mail": "",
                    "phone": ""
                }
            }
        ],
        "CreatedAt": "2024-01-03T15:05:09.503045+03:00",
        "UpdatedAt": "2024-01-03T15:05:09.503045+03:00",
        "DeletedAt": null
    }
  }
  ```

- **shift-schedules/id (sil)**

  - İstek

  ```bash
    curl --location --request DELETE 'http://localhost:9097/shift-scheduler-service/shift-schedules/1'
  ```

  - Yanıt

  ```json
  {
      "status": true,
      "intent": "cld:::shiftservice:::/shift-schedules/:id",
      "message": "Shift Schedule Successfully Deleted"
  }
  ```

- **shift-schedules/id/restore (geri al)**

  - Request

  ```bash
    curl --location --request PATCH 'http://localhost:9097/shift-scheduler-service/shift-schedules/1/restore'
  ```

  - Response

  ```json
  {
      "status": true,
      "intent": "cld:::shiftservice:::/shift-schedules/:id/restore",
      "message": "Shift Schedule Successfully Restored"
  }
  ```

- **shift-schedules (oluştur)**

  - İstek

  ```bash
  curl --location 'http://localhost:9097/shift-scheduler-service/shift-schedules' \
  --header 'Content-Type: application/json' \
  --data '{
      "alias": "Shift 2",
      "description": "Description 1",
      "frequency": 7,
      "start_date": "2023-01-01T00:00:00Z",
      "end_date": "2023-11-01T00:00:00Z",
      "year": 2023,
      "status": 2,
      "organization": [
          {
              "id": 0,
              "name": "Group 2",
              "mail": "",
              "phone": "",
              "description": ""
          }
      ],
      "manager": [
          {
              "id": 0,
              "name": "Manager 2",
              "mail": "",
              "phone": "",
              "description": ""
          }
      ],
      "shifts": [
          {
              "id": 0,
              "start": "2023-01-01 00:00:00",
              "end": "2023-02-01 00:00:00",
              "user": {
                  "name": "User 1",
                  "mail": "",
                  "phone": ""
              }
          },
          {
              "id": 1,
              "start": "2023-01-01 00:00:00",
              "end": "2023-02-01 00:00:00",
              "user": {
                  "name": "User 2",
                  "mail": "",
                  "phone": ""
              }
          },
          {
              "id": 2,
              "start": "2023-01-01 00:00:00",
              "end": "2023-02-01 00:00:00",
              "user": {
                  "name": "User 3",
                  "mail": "",
                  "phone": ""
              }
          }
      ]
  }'
  ```

  - Yanıt

  ```json
  {
    "status": true,
    "intent": "cld:::shiftservice:::/shift-schedules",
    "message": "Shift Schedule Successfully Created"
  }
  ```

- **shift-schedules/id (güncelle)**

  - İstek

  ```bash
  curl --location --request PUT 'http://localhost:9097/shift-scheduler-service/shift-schedules/2' \
  --header 'Content-Type: application/json' \
  --data '{
      "alias": "Shift 2",
      "description": "Description 1",
      "frequency": 7,
      "start_date": "2023-01-01T00:00:00Z",
      "end_date": "2023-11-01T00:00:00Z",
      "year": 2029,
      "status": 2,
      "organization": [
          {
              "id": 0,
              "name": "Group 2",
              "mail": "",
              "phone": "",
              "description": ""
          }
      ],
      "manager": [
          {
              "id": 0,
              "name": "Manager 2",
              "mail": "",
              "phone": "",
              "description": ""
          }
      ],
      "shifts": [
          {
              "id": 0,
              "start": "2023-01-01 00:00:00",
              "end": "2023-02-01 00:00:00",
              "user": {
                  "name": "User 1",
                  "mail": "",
                  "phone": ""
              }
          },
          {
              "id": 1,
              "start": "2023-01-01 00:00:00",
              "end": "2023-02-01 00:00:00",
              "user": {
                  "name": "User 2",
                  "mail": "",
                  "phone": ""
              }
          },
          {
              "id": 2,
              "start": "2023-01-01 00:00:00",
              "end": "2023-02-01 00:00:00",
              "user": {
                  "name": "User 3",
                  "mail": "",
                  "phone": ""
              }
          }
      ]
  }'
  ```

  - Yanıt

  ```json
  {
    "status": true,
    "intent": "cld:::shiftservice:::/shift-schedules/:id",
    "message": "Shift Schedule Successfully Updated"
  }
  ```

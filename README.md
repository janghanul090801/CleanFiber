# FiberEnt | *Clean Architecture in Go* 🎉

<p align='center'>
  <img src='https://res.cloudinary.com/chkilel/image/upload/v1655654392/fiberent/fiberent-preview_lp0p4b.png' alt='FiberEnt' width='60%'/>
</p>


 FiberEnt is a clean architecture implementation in Go with the following frameworks:
- [Fiber](https://github.com/gofiber/fiber) 🚀 is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go.
- [Ent](https://github.com/ent/ent) 🎉 is an entity framework for Go,
Simple, yet powerful ORM for modeling and querying data.

<br/>

## 개발 시작하기
> Docker 필요

Docker 컨테이너 실행
```bash
  make docker-dev # or docker-compose up
```
데이터베이스 마이그레이션 실행
```bash
  make migrate
```
<br />

# 새 엔티티 생성법

Install **Ent** entity framework, check out [https://entgo.io/docs/getting-started#installation](https://entgo.io/docs/getting-started#installation) for more information.

1. 엔티티 스키마 생성

   ```bash
   go run entgo.io/ent/cmd/ent init User # User is the name of the entity
   ```

2. `<project>/ent/schema/user.go` 수정

   - 필드 정의(**[Ent 필드 가이드](https://entgo.io/docs/schema-fields)**)
   - 관계(Edges) 정의(**[Ent Edges 가이드](https://entgo.io/docs/schema-edges)**)

3. 프로젝트 루트에서 코드 생성.

     ```bash
     go generate ./ent
     ```

4. `<project>/entity/user.go` 파일 생성 (엔티티 정의)

5. `<project>/usecase/user` 폴더에서 **Repository 인터페이스**와 **Usecase(Service) 인터페이스** 정의

6. `<project>/usecase/user/service.go` -> Usecase 인터페이스 구현

7. `<project>/infrastructure/ent/repository/user_ent.go` -> Repository 인터페이스 구현.

8. `<project>/api/handler/user.go` 및 `<project>/api/presenter/user.go` 추가 (헨들러 / DTO)

9. `<project>/api/main.go` 파일에 새로운 엔드포인트 등록

## API requests

### Add a user

```
curl -X "POST" "http://localhost:3030/api/v1/users" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
          "email": "adil@mail.com",
          "first_name": "Adil",
          "last_name": "Chehabi",
          "password": "password"
          }'
```
### Update a user

```
curl -X "POST" "http://localhost:3030/api/v1/users/[USER_ID]" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
          "email": "adil@mail.com",
          "first_name": "Adil",
          "last_name": "Chkilel",
          "password": "password"
          }'
```

### Get a user

```
curl "http://localhost:3030/api/v1/users/[USER_ID]" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Delete a user

```
curl -X "DELETE" "http://localhost:3030/api/v1/users/[USER_ID]" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### List all users

```
curl "http://localhost:3030/api/v1/users" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

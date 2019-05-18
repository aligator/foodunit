<?php

namespace foodunit\services;

require_once 'database/Connection.php';

use foodunit\database\Connection;

/**
 * Class OfferService
 * @package foodunit\services
 */
class OfferService
{
    /**
     * @var Connection
     */
    private $db;

    /**
     * OfferService constructor.
     */
    public function __construct()
    {
        $this->db = new Connection();
    }

    /**
     * @return array
     */
    public function getActiveOffers()
    {
        $res = $this->db->query(/** @lang sql */'
            SELECT    id, supplier_id
            FROM      offers
            ORDER BY  id DESC
            LIMIT     1
        ');
        return $res;
    }

    /**
     * @param int $offerId
     * @return array
     */
    public function getAllOrders(int $offerId)
    {
        $bindings = [
            'offer_id' => $offerId
        ];
        $positions = $this->db->query(/** @lang sql */'
            SELECT      o.id, o.offer_id, o.dish_id, s.email, d.name
            FROM        orders o
            INNER JOIN  sessions s
            ON          o.session_id = s.id
            INNER JOIN  dishes d
            ON          o.dish_id = d.id
            WHERE       offer_id = :offer_id
        ', $bindings);

        $orders = [];

        foreach ($positions as $row) {
            $email = $row['email'];

            if (!isset($orders[$email])) {
                $orders[$email] = [
                    'email' => $email,
                    'positions' => []
                ];
            }
            $position = [
                'dish_id' => $row['dish_id'],
                'name' => $row['name']
            ];
            $orders[$email]['positions'][] = $position;
        }
        return $orders;
    }

    /**
     * @param int $offerId
     * @param string $key
     * @return array
     */
    public function getUserOrder(int $offerId, string $key)
    {
        $bindings = [
            'offer_id' => $offerId,
            'key' => $key
        ];
        $res = $this->db->query(/** @lang sql */'
            SELECT      o.dish_id, d.name
            FROM        orders o
            INNER JOIN  dishes d
            ON          o.dish_id = d.id
            WHERE       o.offer_id = :offer_id
            AND         o.session_id = (
                SELECT  s.id
                FROM    sessions s
                WHERE   s.key = :key
                LIMIT   1
            )
        ', $bindings);

        return $res;
    }

    /**
     * @param int $offerId
     * @param int $dishId
     * @param string $key
     * @return bool
     */
    public function addDishToOrder(int $offerId, int $dishId, string $key)
    {
        $bindings = [
            'offer_id' => $offerId,
            'dish_id' => $dishId,
            'key' => $key
        ];
        $success = $this->db->exec(/** @lang sql */'
            INSERT INTO orders (offer_id, dish_id, session_id)
            SELECT      :offer_id, :dish_id, id
            FROM        sessions s
            WHERE       s.key = :key
            LIMIT       1
        ', $bindings);

        return $success;
    }

    /**
     * @param int $offerId
     * @param int $dishId
     * @param string $key
     * @return bool
     */
    public function deleteDishFromOrder(int $offerId, int $dishId, string $key)
    {
        $bindings = [
            'offer_id' => $offerId,
            'dish_id' => $dishId,
            'key' => $key
        ];
        $success = $this->db->exec(/** @lang sql */'
            DELETE FROM orders
            WHERE       offer_id = :offer_id
            AND         dish_id = :dish_id
            AND         session_id = (
                SELECT  s.id
                FROM    sessions s
                WHERE   s.key = :key
                LIMIT   1
            )
        ', $bindings);

        return $success;
    }

    /**
     * @param int $offerId
     * @param string $key
     * @return string
     */
    public function getRemark(int $offerId, string $key)
    {
        $bindings = [
            'offer_id' => $offerId,
            'key' => $key
        ];
        $remark = $this->db->query(/** @lang sql */'
            SELECT      r.remark
            FROM        remarks r
            WHERE       r.offer_id = :offer_id
            AND         r.session_id = (
                SELECT  s.id
                FROM    sessions s
                WHERE   s.key = :key
                LIMIT   1
            )
            ORDER BY    r.id DESC
            LIMIT       1
        ', $bindings);

        $remark = $remark[0]['remark'];

        return $remark;
    }

    /**
     * @param int $offerId
     * @param string $remark
     * @param string $key
     * @return bool
     */
    public function insertRemark(int $offerId, string $remark, string $key)
    {
        $bindings = [
            'offer_id' => $offerId,
            'remark' => $remark,
            'key' => $key
        ];
        $success = $this->db->exec(/** @lang sql */'
            INSERT INTO remarks (offer_id, session_id, remark)
            SELECT      :offer_id, id, :remark
            FROM        sessions s
            WHERE       s.key = :key
            LIMIT       1
        ', $bindings);

        return $success;
    }
}

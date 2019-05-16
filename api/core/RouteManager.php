<?php

namespace foodunit\core;

require_once 'RouteHandler.php';

/**
 * Class RouteManager
 * @package foodunit\core
 */
class RouteManager
{
    /**
     * @var array
     */
    private $mappings;

    /**
     * RouteManager constructor.
     */
    public function __construct()
    {
        $this->mappings = $this->loadMappings();
    }

    /**
     * @return array
     */
    public function mappings()
    {
        return $this->mappings;
    }

    /**
     * @return array
     */
    private function loadMappings()
    {
        return [
            '/offers'                       => 'foodunit\core\RouteHandler:offers',
            '/dishes/{supplier}'            => 'foodunit\core\RouteHandler:dishes',
            '/orders/{offer}'               => 'foodunit\core\RouteHandler:orders',
            '/user-order/{offer}'           => 'foodunit\core\RouteHandler:userOrder',
            '/add/{offer}/{dish}'           => 'foodunit\core\RouteHandler:add',
            '/del/{offer}/{dish}'           => 'foodunit\core\RouteHandler:del',
            '/remark/{offer}'               => 'foodunit\core\RouteHandler:remark'
        ];
    }
}